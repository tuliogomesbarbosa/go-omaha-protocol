package omaha

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

const StableChannel = "https://stable.release.core-os.net/amd64-usr/"

// Data : Each <data> tag in the request represents either a request for additional textual information from the server, or provides additional textual information to the server.
type Data struct {
    XMLName         xml.Name    `xml:"data"`
    Name            string      `xml:"name,attr,omitempty"`
    Index           int         `xml:"index,attr,omitempty"`
}

// Disabled : an integral reason that the app is disabled.
type Disabled struct {
    XMLName         xml.Name    `xml:"disabled"` 
    Reason          string      `xml:"reason,attr,omitempty"`
}

// Packages : A <packages> tag simply contains several <package>s
type Packages struct {
    XMLName         xml.Name    `xml:"packages"`
    Package         []Package   `xml:"package"`
}

// Package : A <package> tag gives information about an installed package.
type Package struct {
    XMLName         xml.Name    `xml:"package"`
    Fingerprint     string      `xml:"fingerprint,attr,omitempty"`
}

// Ping : Any <ping>s contained in a request are used to count active users and potentially deduplicate requests from the same client.
type Ping struct {
    XMLName         xml.Name    `xml:"ping"`
    Active          string      `xml:"active,attr,omitempty"`
    A               uint        `xml:"a,attr,omitempty"`
    R               uint        `xml:"r,attr,omitempty"`
    AD              uint        `xml:"ad,attr,omitempty"`
    RD              uint        `xml:"rd,attr,omitempty"`
    PingFreshness   int64       `xml:"ping_freshness,attr,omitempty"`
    Status          string      `xml:"status,attr"`
}

// UpdateCheck : 
type UpdateCheck struct {
	XMLName             xml.Name    `xml:"updatecheck"`
    TTToken             string      `xml:"tttoken,attr,omitempty"`
    UpdateDisable       bool        `xml:"updatedisabled,attr,omitempty"`
    TargetVersionPrefix string      `xml:"targetversionprefix,attr,omitempty"`
	//Urls                *Urls     `xml:"urls"`
	//Manifest            *Manifest `xml:"manifest"`	
	Status              string      `xml:"status,attr"`
}

// Event : Throughout and at the end of an update flow, the client MAY send event reports by sending one or more requests containing an <event>.
type Event struct {
// <event>s should never appear in the same request as an <updatecheck>.
    XMLName                     xml.Name    `xml:"event"`
    EventType                   int         `xml:"eventtype,attr"`
    EventResult                 int         `xml:"eventresult,attr,omitempty"`
    ErrorCode                   int         `xml:"errorcode,attr,omitempty"`
    ExtraCode1                  int         `xml:"extracode1,attr,omitempty"`
    ErrorCat                    int         `xml:"errorcat,attr,omitempty"`
    DownloadTimeMs              int         `xml:"download_time_ms,attr,omitempty"`
    Downloaded                  int         `xml:"downloaded,attr,omitempty"`
    Total                       int         `xml:"total,attr,omitempty"`
    UpdateCheckTimeMs           int         `xml:"update_check_time_ms,attr,omitempty"`
    InstallTimeMs               int         `xml:"install_time_ms,attr,omitempty"`
    SourceURLIndex              string      `xml:"source_url_index,attr"`
    StateCancelled              int         `xml:"state_cancelled,attr,omitempty"`
    TimeSinceUpdateAvailableMs  int         `xml:"time_since_update_available_ms,attr,omitempty"`
    URL                         string      `xml:"url,attr,omitempty"`
    NextVersion                 string      `xml:"nextversion,attr,omitempty"`
    PreviousVersion             string      `xml:"previousversion,attr,omitempty"`
    NextFP                      string      `xml:"nextfp,attr,omitempty"`
    PreviousFP                  string      `xml:"previousfp,attr,omitempty"`
}

// App : Each product that is contained in the request is represented by exactly one <app> tag.
type App struct {
    XMLName         xml.Name    `xml:"app"`
    Data            []Data      `xml:"data,omitempty"`
    Disabled        []Disabled  `xml:"disabled,omitempty"`
    Packages        Packages    `xml:"packages,omitempty"`
    Ping            Ping        `xml:"ping"`
    //at most one event or UpdateCheck
    Events          []Event     `xml:"event,omitempty"` //one or more
    UpdateCheck     UpdateCheck `xml:"updatecheck,omitempty"` //exactly one
    AppID           string      `xml:"appid,attr"`
    Version         string      `xml:"version,attr,omitempty"`
    Lang            string      `xml:"lang,attr,omitempty"`
    Brand           string      `xml:"brand,attr,omitempty"`
    Client          string      `xml:"client,attr,omitempty"`
    Enabled         string      `xml:"enabled,attr,omitempty"`
    Experiments     string      `xml:"experiments,attr,omitempty"`
    IID             string      `xml:"iid,attr"`
    InstallAge      string      `xml:"installage,attr,omitempty"`
    InstallDate     string      `xml:"installdate,attr,omitempty"`
    InstallSource   string      `xml:"installsource,attr,omitempty"`
    IsMachine       string      `xml:"ismachine,attr,omitempty"`
    Track           string      `xml:"track,attr,omitempty"` //TAG equivalent
    Fingerprint     string      `xml:"fingerprint,attr,omitempty"`
    COHORT          string      `xml:"cohort,attr,omitempty"`
    COHORTHint      string      `xml:"cohorthint,attr,omitempty"`
    COHORTName      string      `xml:"cohortname,attr,omitempty"`
    Status          string      `xml:"status,attr,omitempty"`
}

// Os : 
type Os struct {
    XMLName         xml.Name    `xm:"os"`
    Platform        string      `xml:"platform,attr,omitempty"`
	Version         string      `xml:"version,attr,omitempty"`
	Sp              string      `xml:"sp,attr,omitempty"`
	Arch            string      `xml:"arch,attr,omitempty"`
}

// Hw : 
type Hw struct {
    XMLName         xml.Name    `xml:"hw"`
    Sse             string      `xml:"sse,attr,omitempty"`
    Sse2            string      `xml:"sse2,attr,omitempty"`
    Sse3            string      `xml:"sse3,attr,omitempty"`
    Sse41           string      `xml:"sse41,attr,omitempty"`
    Sse42           string      `xml:"sse42,attr,omitempty"`
    Ssse3           string      `xml:"ssse3,attr,omitempty"`
    Avx             string      `xml:"avx,attr,omitempty"`
    Physmemory      string      `xml:"physmemory,attr,omitempty"`
}

// Request : 
type Request struct {
    XMLName         xml.Name    `xml:"request"`
    Apps            []*App      `xml:"app"`
    Hw              Hw          `xml:"hw"`
    Os              Os          `xml:"os"`
    Dedup           string      `xml:"dedup,attr,omitempty"`
    DlPref          string      `xml:"dlpref,attr,omitempty"`
    InstallSource   string      `xml:"installsource,attr,omitempty"`
    IsMachine       string      `xml:"ismachine,attr,omitempty"`
    OriginURL       string      `xml:"originurl,attr,omitempty"`
    Protocol        string      `xml:"protocol,attr"`
    RequestID       string      `xml:"requestid,attr,omitempty"`
    SessionID       string      `xml:"sessionid,attr,omitempty"`
    TestSource      string      `xml:"testsource,attr,omitempty"`
    UpdaterChannel  string      `xml:"updaterchannel,attr,omitempty"`
    UserID          string      `xml:"userid,attr,omitempty"`
    Version         string      `xml:"version,attr,omitempty"`
}

type DayStart struct {   
    ElapsedSeconds  int         `xml:"elapsed_seconds,attr"`
    ElapsedDays     int         `xml:"elapsed_days,attr"`
}

// Response :
type Response struct {
    XMLName         xml.Name    `xml:"response"`
    Protocol        string      `xml:"protocol,attr"`
    Server          string      `xml:"server,attr"`
    Apps            []*App      `xml:"app,omitempty"`
    DayStart        DayStart    `xml:"daystart,omitempty"`       
}

// EventTypes : 
var EventTypes = map[int]string{
	0:   "unknown",
	1:   "download complete",
	2:   "install complete",
	3:   "update complete",
	4:   "uninstall",
	5:   "download started",
	6:   "install started",
	9:   "new application install started",
	10:  "setup started",
	11:  "setup finished",
	12:  "update application started",
	13:  "update download started",
	14:  "update download finished",
	15:  "update installer started",
	16:  "setup update begin",
	17:  "setup update complete",
	20:  "register product complete",
	30:  "OEM install first check",
	40:  "app-specific command started",
	41:  "app-specific command ended",
    50:  "update-check failure",    
	100: "setup failure",
	102: "COM server failure",
	103: "setup update failure",
}

// EventResults :
var EventResults = map[int]string{
	0:  "error",
	1:  "success",
	2:  "success reboot",
	3:  "success restart browser",
	4:  "cancelled",
	5:  "error installer MSI",
	6:  "error installer other",
	7:  "noupdate",
	8:  "error installer system",
	9:  "update deferred",
	10: "handoff error",
}

// StatesCancelled : 
var StatesCancelled = map[int]string{
    0:  "unknown or not-cancelled",
    1:  "initializing",
    2:  "waiting to check for update",
    3:  "checking for update",
    4:  "update available",
    5:  "waiting to download",
    6:  "retrying download",
    7:  "downloading",
    8:  "download complete",
    9:  "extracting",
    10: "applying differential patch",
    11: "ready to install",
    12: "waiting to install",
    13: "installing",
    14: "install complete",
    15: "paused",
    16: "no update",
    17: "error",
}

func NewApp(appid string) *App {
    return &App{AppID: appid}
}

func (r *Response) AddApp(appid string) {
    a := NewApp(appid)
    r.Apps = append(r.Apps, a)
}

// NewNoUpdate :
func NewNoUpdate(app *App) *Response {
    r := new(Response)
    
    return r
} 

func CheckForUpdate(r *Request) *Response {
    return NewNoUpdate(r.Apps[0])
}

func main()  {
    //testing omaha requests parsing
    file, err := os.Open("data/no-update/request.xml")
    if err != nil {
        return
    }
    request, err := ioutil.ReadAll(file)
    if err != nil {
        return
    }
    r := Request{}
    if err = xml.Unmarshal(request, &r); err != nil {
        return
    }
    switch {
    case r.Protocol == "":
        fmt.Println("Protocol should be provided.")
        return
    case r.Os.Version != "Ladybug":
        fmt.Println("Ladybug is the only supported Os.")
        return    
    }

    response := CheckForUpdate(&r);
    enc := xml.NewEncoder(os.Stdout)
    enc.Indent("  ", "    ")
    if err := enc.Encode(response); err != nil {
        fmt.Printf("error: %v\n", err)
    }
    
}