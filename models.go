package insteon

type Command struct {
	Device_Id     int    `json:"device_id,omitempty"`
	Scene_Id      int    `json:"scene_id,omitempty"`
	House_Id      int    `json:"house_id,omitempty"`
	Command       string `json:"command"`
	Level         int    `json:"level,omitempty"`
	Group         int    `json:"group,omitempty"`
	State         int    `json:"state,omitempty"`
	Flag          string `json:"flag,omitempty"`
	Temp          int    `json:"temp,omitempty"`
	As_Controller bool   `json:"as_controller,omitempty"`
}
type CommandResponse struct {
	Id       int                    `json:"id"`
	Status   string                 `json:"status"`
	Link     string                 `json:"link"`
	Command  Command                `json:"command"`
	Response map[string]interface{} `json:"response"`
}

type CameraResponse struct {
	CameraList []Camera
}
type Camera struct {
	HouseID          int
	CameraID         int
	CameraName       string
	IP               string
	WebPort          int
	MediaPort        int
	Gateway          string
	Url              string
	Username         string
	Password         string
	Model            int
	DeviceType       int
	StreamType       int
	UniqueIdentifier string
	MacAddress       string
	SystemVersion    string
	AppVersion       string
	Favorite         bool
	SerialNumber     string
	DhcpEnabled      bool `json:"dhcpEnabled"`
}
type RoomResponse struct {
	RoomList []Room
}
type Room struct {
	HouseID       int
	RoomID        int
	RoomName      string
	Visible       bool
	Favorite      bool
	DefaultCamera int
	CameraList    []Camera
	DeviceList    []Device
	SceneList     []Scene
}
type SceneResponse struct {
	SceneList []Scene
}
type Scene struct {
	HouseID         int
	SceneID         int
	SceneName       string
	StatusDevice    string
	OnTime          string
	OffTime         string
	CustomOn        string
	CustomOff       string
	Group           int
	Visible         bool
	Favorite        bool
	AutoStatus      bool
	DayMask         int
	TimerEnabled    bool
	EnableCustomOn  bool
	EnableCustomOff bool
	DeviceList      []Device
}
type DeviceResponse struct {
	DeviceList []Device
}
type Device struct {
	HouseID               int
	DeviceID              int
	DeviceName            string
	InsteonID             string
	FirmwareVersion       int
	OnTime                string
	OffTime               string
	TimerEnabled          bool
	Group                 int
	DeviceType            int
	DevCat                int
	SubCat                int
	AutoStatus            bool
	CustomOn              string
	CustomOff             string
	EnableCustomOn        bool
	EnableCustomOff       bool
	DimLevel              int
	RampRate              int
	OperationFlags        int
	LEDLevel              int
	AlertsEnabled         bool
	AlertOn               int
	AlertOff              int
	Favorite              bool
	Humidity              bool
	DayMask               int
	LinkWithHub           int
	BeepOnPress           bool
	LocalProgramLock      bool
	BlinkOnTraffic        bool
	ErrorBlink            bool
	ConfiguredGroups      int
	InsteonEngine         int
	SerialNumber          string
	Manufacturer          string
	ProductType           string
	User                  string
	UserID                string
	AccessToken           string
	AccessTokenExpiration string
	GroupList             []Group
}
type Group struct {
	GroupNum   int
	GroupName  string
	GroupState int
	SceneID    int
}

type BearerResponse struct {
	Refresh_Token string `json:"refresh_token"`
	Access_Token  string `json:"access_token"`
	Token_Type    string `json:"token_type"`
	Expires_In    int    `json:"expires_in"`
}

type CommandStatusResponse struct {
	Status string `json:"status"`
	Link   string `json:"link"`
	ID     int    `json:"id"`
}

type Authorization struct {
	ClientID     string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}
