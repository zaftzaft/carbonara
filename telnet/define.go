/*
	ref: https://github.com/wireshark/wireshark/blob/master/epan/dissectors/packet-telnet.c
*/
package telnet

type TelnetCmd uint8

const (
	IAC  TelnetCmd = 255
	DONT TelnetCmd = 254
	DO   TelnetCmd = 253
	WONT TelnetCmd = 252
	WILL TelnetCmd = 251
	SB   TelnetCmd = 250
	SE   TelnetCmd = 240
)

func (c TelnetCmd) String() string {
	switch c {
	case IAC:
		return "IAC"
	case DONT:
		return "DONT"
	case DO:
		return "DO"
	case WONT:
		return "WONT"
	case WILL:
		return "WILL"
	case SB:
		return "SB"
	case SE:
		return "SE"
	default:
		return ""
	}
}

type TelnetOp uint8

const (
	BinaryTransmission              TelnetOp = iota /*RFC856*/
	Echo                            TelnetOp = iota /*RFC857*/
	Reconnection                    TelnetOp = iota /*DODProtocolHandbook*/
	SuppressGoAhead                 TelnetOp = iota /*RFC858*/
	ApproxMessageSizeNegotiation    TelnetOp = iota /*Ethernetspec(!)*/
	Status                          TelnetOp = iota /*RFC859*/
	TimingMark                      TelnetOp = iota /*RFC860*/
	RemoteControlledTransandEcho    TelnetOp = iota /*RFC726*/
	OutputLineWidth                 TelnetOp = iota /*DODProtocolHandbook*/
	OutputPageSize                  TelnetOp = iota /*DODProtocolHandbook*/
	OutputCarriageReturnDisposition TelnetOp = iota /*RFC652*/
	OutputHorizontalTabStops        TelnetOp = iota /*RFC653*/
	OutputHorizontalTabDisposition  TelnetOp = iota /*RFC654*/
	OutputFormfeedDisposition       TelnetOp = iota /*RFC655*/
	OutputVerticalTabstops          TelnetOp = iota /*RFC656*/
	OutputVerticalTabDisposition    TelnetOp = iota /*RFC657*/
	OutputLinefeedDisposition       TelnetOp = iota /*RFC658*/
	ExtendedASCII                   TelnetOp = iota /*RFC698*/
	Logout                          TelnetOp = iota /*RFC727*/
	ByteMacro                       TelnetOp = iota /*RFC735*/
	DataEntryTerminal               TelnetOp = iota /*RFC732,RFC1043*/
	SUPDUP                          TelnetOp = iota /*RFC734,RFC736*/
	SUPDUPOutput                    TelnetOp = iota /*RFC749*/
	SendLocation                    TelnetOp = iota /*RFC779*/
	TerminalType                    TelnetOp = iota /*RFC1091*/
	EndofRecord                     TelnetOp = iota /*RFC885*/
	TACACSUserIdentification        TelnetOp = iota /*RFC927*/
	OutputMarking                   TelnetOp = iota /*RFC933*/
	TerminalLocationNumber          TelnetOp = iota /*RFC946*/
	Telnet3270Regime                TelnetOp = iota /*RFC1041*/
	X3PAD                           TelnetOp = iota /*RFC1053*/
	NegotiateAboutWindowSize        TelnetOp = iota /*RFC1073,DW183*/
	TerminalSpeed                   TelnetOp = iota /*RFC1079*/
	RemoteFlowControl               TelnetOp = iota /*RFC1372*/
	Linemode                        TelnetOp = iota /*RFC1184*/
	XDisplayLocation                TelnetOp = iota /*RFC1096*/
	EnvironmentOption               TelnetOp = iota /*RFC1408,RFC1571*/
	AuthenticationOption            TelnetOp = iota /*RFC2941*/
	EncryptionOption                TelnetOp = iota /*RFC2946*/
	NewEnvironmentOption            TelnetOp = iota /*RFC1572*/
	TN3270E                         TelnetOp = iota /*RFC1647*/
	XAUTH                           TelnetOp = iota /*XAUTH*/
	CHARSET                         TelnetOp = iota /*CHARSET*/
	RemoteSerialPort                TelnetOp = iota /*RemoteSerialPort*/
	COMPortControl                  TelnetOp = iota /*RFC2217*/
	SuppressLocalEcho               TelnetOp = iota /*draft-rfced-exp-atmar-00*/
	StartTLS                        TelnetOp = iota /*draft-ietf-tn3270e-telnet-tls-06*/
	KERMIT                          TelnetOp = iota /*RFC2840*/
	SENDURL                         TelnetOp = iota /*draft-croft-telnet-url-trans-00*/
	FORWARD_X                       TelnetOp = iota /*draft-altman-telnet-fwdx-03*/
)

func (o TelnetOp) String() string {
	switch o {
	case BinaryTransmission:
		return "Binary Transmission"
	case Echo:
		return "Echo"
	case Reconnection:
		return "Reconnection"
	case SuppressGoAhead:
		return "Suppress Go Ahead"
	case ApproxMessageSizeNegotiation:
		return "Approx Message Size Negotiation"
	case Status:
		return "Status"
	case TimingMark:
		return "Timing Mark"
	case RemoteControlledTransandEcho:
		return "Remote Controlled Trans and Echo"
	case OutputLineWidth:
		return "Output Line Width"
	case OutputPageSize:
		return "Output Page Size"
	case OutputCarriageReturnDisposition:
		return "Output Carriage-Return Disposition"
	case OutputHorizontalTabStops:
		return "Output Horizontal Tab Stops"
	case OutputHorizontalTabDisposition:
		return "Output Horizontal Tab Disposition"
	case OutputFormfeedDisposition:
		return "Output Formfeed Disposition"
	case OutputVerticalTabstops:
		return "Output Vertical Tabstops"
	case OutputVerticalTabDisposition:
		return "Output Vertical Tab Disposition"
	case OutputLinefeedDisposition:
		return "Output Linefeed Disposition"
	case ExtendedASCII:
		return "Extended ASCII"
	case Logout:
		return "Logout"
	case ByteMacro:
		return "Byte Macro"
	case DataEntryTerminal:
		return "Data Entry Terminal"
	case SUPDUP:
		return "SUPDUP"
	case SUPDUPOutput:
		return "SUPDUP Output"
	case SendLocation:
		return "Send Location"
	case TerminalType:
		return "Terminal Type"
	case EndofRecord:
		return "End of Record"
	case TACACSUserIdentification:
		return "TACACS User Identification"
	case OutputMarking:
		return "Output Marking"
	case TerminalLocationNumber:
		return "Terminal Location Number"
	case Telnet3270Regime:
		return "Telnet 3270 Regime"
	case X3PAD:
		return "X.3 PAD"
	case NegotiateAboutWindowSize:
		return "Negotiate About Window Size"
	case TerminalSpeed:
		return "Terminal Speed"
	case RemoteFlowControl:
		return "Remote Flow Control"
	case Linemode:
		return "Linemode"
	case XDisplayLocation:
		return "X Display Location"
	case EnvironmentOption:
		return "Environment Option"
	case AuthenticationOption:
		return "Authentication Option"
	case EncryptionOption:
		return "Encryption Option"
	case NewEnvironmentOption:
		return "New Environment Option"
	case TN3270E:
		return "TN3270E"
	case XAUTH:
		return "XAUTH"
	case CHARSET:
		return "CHARSET"
	case RemoteSerialPort:
		return "Remote Serial Port"
	case COMPortControl:
		return "COM Port Control"
	case SuppressLocalEcho:
		return "Suppress Local Echo"
	case StartTLS:
		return "Start TLS"
	case KERMIT:
		return "KERMIT"
	case SENDURL:
		return "SEND-URL"
	case FORWARD_X:
		return "FORWARD_X"
	}

	return ""
}
