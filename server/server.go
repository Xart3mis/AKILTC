package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/Xart3mis/GoHkarComms/client_data_pb"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/magodo/textinput"

	"google.golang.org/grpc"
)

type server struct {
	pb.ConsumerServer
}

var current_id string = ""

var client_ids []string
var client_mapids map[int]string = make(map[int]string)

var client_onscreentext map[string]string = make(map[string]string)
var client_execcommand map[string]string = make(map[string]string)
var client_execoutput map[string]string = make(map[string]string)

var prev_msg string = ""

func main() {
	go func() {
		p := tea.NewProgram(initialModel(), tea.WithAltScreen())

		if err := p.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterConsumerServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func (s *server) GetCommand(ctx context.Context, cid *pb.ClientDataRequest) (*pb.ClientExecData, error) {
	if !Contains(client_ids, cid.ClientId) {
		client_ids = append(client_ids, cid.ClientId)
	}

	if cid.ClientId == current_id && cid != nil {
		x := client_execcommand[current_id]
		client_execcommand[current_id] = ""
		return &pb.ClientExecData{
			ShouldExec: len(x) > 0,
			Command:    x}, nil
	}
	return &pb.ClientExecData{}, nil
}

func (s *server) SetCommandOutput(ctx context.Context, in *pb.ClientExecOutput) (*pb.Void, error) {
	if id := in.GetId(); id != nil {
		if id.ClientId == current_id {
			client_execoutput[current_id] = string(in.GetOutput())
		}
	}

	return &pb.Void{}, nil
}

func (s *server) SubscribeOnScreenText(r *pb.ClientDataRequest, in pb.Consumer_SubscribeOnScreenTextServer) error {
	if !Contains(client_ids, r.ClientId) {
		client_ids = append(client_ids, r.ClientId)
	}
	for {
		if r.GetClientId() == current_id {
			in.Send(&pb.ClientDataOnScreenTextResponse{OnScreen: &pb.ClientOnScreenData{
				ShouldUpdate: len(client_onscreentext[current_id]) > 0,
				OnScreenText: client_onscreentext[current_id]}})
		} else {
			in.Send(&pb.ClientDataOnScreenTextResponse{OnScreen: &pb.ClientOnScreenData{
				ShouldUpdate: false,
				OnScreenText: ""}})
		}
	}
}

type model struct {
	textInput           textinput.Model
	typing              bool
	showhelplist        bool
	showclientlist      bool
	showok              bool
	shownotvalidcid     bool
	shownotvalidtext    bool
	shownotvalidcommand bool
	showexecout         bool
	err                 error
	clients             []string
}

func initialModel() model {
	ti := textinput.NewModel()
	ti.Placeholder = "help"

	red := color.New(color.FgRed).SprintFunc()
	ti.Prompt = red(">> ")
	ti.PromptStyle.Bold(true)

	ti.PromptStyle.PaddingRight(10)
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	ti.CandidateWords = []string{"help", "settext", "select", "exit", "exec", "list_clients"}
	ti.CandidateViewMode = textinput.CandidateViewHorizental

	return model{
		textInput:           ti,
		err:                 nil,
		typing:              true,
		showhelplist:        false,
		showok:              false,
		showclientlist:      false,
		shownotvalidcid:     false,
		shownotvalidtext:    false,
		shownotvalidcommand: false,
		showexecout:         false,
		clients:             client_ids,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	for idx, client := range client_ids {
		client_mapids[idx] = client
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			os.Exit(0)
			return m, tea.Quit

		case "esc", "q":
			m.shownotvalidcommand = false
			m.shownotvalidtext = false
			m.shownotvalidcid = false
			m.showclientlist = false
			m.showhelplist = false
			m.showok = false
			m.showexecout = false
			return m, nil

		case "enter":
			m.clients = client_ids

			if len(m.textInput.Value()) > 4 && m.textInput.Value()[:4] == "exec" {
				split_str := strings.Split(m.textInput.Value(), " ")
				if len(split_str) <= 1 {
					m.shownotvalidcommand = true
					return m, nil
				}
				client_execcommand[current_id] = strings.Join(split_str[1:], " ")
				m.showexecout = true
			}

			if len(m.textInput.Value()) > 6 && m.textInput.Value()[:6] == "select" {
				split_str := strings.Split(m.textInput.Value(), " ")
				if len(split_str) > 2 {
					panic("select only takes 1 argument")
				}
				val, err := strconv.Atoi(split_str[1])
				if err != nil {
					m.shownotvalidcid = true
					return m, nil
				}

				if Contains(client_ids, client_mapids[val]) {
					m.showok = true
					current_id = client_mapids[val]
					return m, nil
				}

				m.shownotvalidcid = true
			}

			if len(m.textInput.Value()) > 7 && m.textInput.Value()[:7] == "settext" {
				split_str := strings.Split(m.textInput.Value(), " ")
				if len(split_str) <= 1 {
					m.shownotvalidtext = true
					return m, nil
				} /*else if len(client_mapids) < 1 {
					//
				}*/
				client_onscreentext[current_id] = strings.Join(split_str[1:], " ")
				m.showok = true
			}

			switch m.textInput.Value() {
			case "help":
				m.showhelplist = true
				return m, nil
			case "exit":
				return m, tea.Quit
			case "list_clients":
				m.showclientlist = true
				return m, nil
			default:
				return m, nil
			}
		}
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd

}

func (m model) View() string {
	if m.showhelplist {
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		return m.textInput.View() + "\n\n" + green("help") + "\nshow this help dialog (usage: " + yellow("help") + ")\n\n" +
			green("select") + "\nselect a client to connect to (usage: " + yellow("select [`client id`]") + ")\n\n" +
			green("exit") + "\nexit the server (usage: " + yellow("exit") + ")\n\n" +
			green("settext") + "\nset on screen text for selected client (usage: " + yellow("settext [`text`]") + ")\n\n" +
			green("exec") + "\nexecute command on selected client (usage: " + yellow("exec [`command string`]") + ")\n\n" +
			green("list_clients") + "\nlist currently connected clients (usage: " + yellow("list_clients [client id]") + ")\n"
	}
	if m.showclientlist {
		Magenta := color.New(color.FgMagenta).SprintFunc()
		b, _ := json.MarshalIndent(client_mapids, "", "\t")
		return m.textInput.View() + "\n\n" + Magenta(string(b))
	}
	if m.shownotvalidcid {
		red := color.New(color.FgRed).SprintFunc()
		return m.textInput.View() + "\n\n" + red("Not a valid client ID.")
	}
	if m.shownotvalidcommand {
		red := color.New(color.FgRed).SprintFunc()
		return m.textInput.View() + "\n\n" + red("Not a valid command.")
	}
	if m.showok {
		Green := color.New(color.FgGreen).SprintFunc()
		return m.textInput.View() + "\n\n" + Green("OK.")
	}
	if m.shownotvalidtext {
		red := color.New(color.FgRed).SprintFunc()
		return m.textInput.View() + "\n\n" + red("settext takes multiple arguments.")
	}
	if m.showexecout {
		x := client_execoutput[current_id]
		client_execoutput[current_id] = ""
		return m.textInput.View() + "\n\n" + x
	}

	return m.textInput.View()
}
