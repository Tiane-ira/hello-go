package fsm

import "testing"

var data = `Routing Engine status:
  Slot 0:
    Current state                  Master
    Election priority              Master (default)
    Temperature                 39 degrees C / 102 degrees F
    CPU temperature             55 degrees C / 131 degrees F
    DRAM                      2048 MB
    Memory utilization          76 percent
    CPU utilization:
      User                      95 percent
      Background                 0 percent
      Kernel                     4 percent
      Interrupt                  1 percent
      Idle                       0 percent
    Model                          RE-4.0
    Serial ID                      xxxxxxxxxxxx
    Start time                     2008-04-10 20:32:25 PDT
    Uptime                         180 days, 22 hours, 45 minutes, 20 seconds
    Load averages:                 1 minute   5 minute  15 minute
                                       0.96       1.03       1.03
Routing Engine status:
  Slot 1:
    Current state                  Backup
    Election priority              Backup
    Temperature                 30 degrees C / 86 degrees F
    CPU temperature             31 degrees C / 87 degrees F
    DRAM                      2048 MB
    Memory utilization          14 percent
    CPU utilization:
      User                       0 percent
      Background                 0 percent
      Kernel                     0 percent
      Interrupt                  1 percent
      Idle                      99 percent
    Model                          RE-4.0
    Serial ID                      xxxxxxxxxxxx
    Start time                     2008-01-22 07:32:10 PST
    Uptime                         260 days, 10 hours, 45 minutes, 39 seconds
`

var template = `Value Required Slot (\d+)
Value State (\w+)
Value Temp (\d+)
Value CPUTemp (\d+)
Value DRAM (\d+)
Value Model (\S+)

Start
  ^Routing Engine status: -> Record
  ^\s+Slot\s+${Slot}
  ^\s+Current state\s+${State}
  ^\s+Temperature\s+${Temp} degrees
  ^\s+CPU temperature\s+${CPUTemp} degrees
  ^\s+DRAM\s+${DRAM} MB
  ^\s+Model\s+${Model} -> Start
`

func TestParseFsm(t *testing.T) {
	result, err := ParseFsm(data, template)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

var data2 = `
Interface	Status	Link	ActSpeed	ActDuplex VSL	Type	Pvid	Desc
---------------------------------------------------------------------------
100ge0/1 	Disabled	1	 Unknown	Unknown		Nni			2 
100ge0/2	Enabled		3	Unknown	Unknown		Nni			2
`

var template2 = `Value Port (\S+ge\S+)
Value Status (\S+)
Value List Link (\d+)

Start
 ^${Port}\s+${Status}\s+${Link} -> Record
`

func TestParseFsm2(t *testing.T) {
	result, err := ParseFsm(data2, template2)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

var data3 = `
100ge0/1
State: Disable1

100ge0/2
State: Disable2

100ge0/3
State: Disable3
`

var template3 = `Value Port (\S+ge\S+)
Value Status (\S+)

Start
 ^\S+ge\S+ -> Continue.Record
 ^${Port}
 ^State: ${Status}
`

func TestParseFsm3(t *testing.T) {
	result, err := ParseFsm(data3, template3)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
