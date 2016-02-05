package ripego

import (
	"log"
	_ "log"
	"testing"
)

func TestTcpContent(t *testing.T) {

	d, err := GetTcpContent("178.18.196.250", "whois.ripe.net")

	if err != nil {
		t.Fatal(err.Error())
	}

	if d == "" {
		t.Fatal("TCP data obtained")
	}

	log.Println(d)
}

func TestRipeWhoisData(t *testing.T) {

	RSPDATA := `% This is the RIPE Database query service.
% The objects are in RPSL format.
%
% The RIPE Database is subject to Terms and Conditions.
% See http://www.ripe.net/db/support/db-terms-conditions.pdf

% Note: this output has been filtered.
%       To receive output for a database update, use the "-B" flag.

% Information related to '178.18.192.0 - 178.18.207.255'

% Abuse contact for '178.18.192.0 - 178.18.207.255' is 'noc@vt.com.tr'

inetnum:        178.18.192.0 - 178.18.207.255
netname:        TR-VARGONEN-20100423
descr:          Vargonen Teknoloji ve Bilisim Sanayi Ticaret Anonim Sirketi
country:        TR
org:            ORG-VTIv1-RIPE
admin-c:        VT5050-RIPE
tech-c:         VT5050-RIPE
status:         ALLOCATED PA
mnt-by:         RIPE-NCC-HM-MNT
mnt-lower:      MNT-VRGN
mnt-routes:     MNT-VRGN
created:        2010-04-23T15:52:35Z
last-modified:  2014-12-11T10:31:35Z
source:         RIPE # Filtered

organisation:   ORG-VTIv1-RIPE
org-name:       Vargonen Teknoloji ve Bilisim Sanayi Ticaret Anonim Sirketi
org-type:       LIR
address:        Vargonen Teknoloji ve Bilisim Sanayi Ticaret Anonim Sirketi
address:        Ankara Cad No 81 Bayrakli Tower Kat 19 Bayrakli
address:        35030
address:        IZMIR
address:        TURKEY
phone:          +908506600099
fax-no:         +902324570021
abuse-c:        AR17405-RIPE
mnt-ref:        MNT-VRGN
mnt-ref:        RIPE-NCC-HM-MNT
mnt-by:         RIPE-NCC-HM-MNT
abuse-mailbox:  noc@vargonen.com
created:        2010-04-07T13:50:24Z
last-modified:  2015-01-07T21:19:36Z
source:         RIPE # Filtered

person:         Vargonen LIR Admin
address:        Ankara Cad. No:81 Bayrakli Tower
address:        Kat:19
address:        Bayrakli - Izmir
address:        TR
phone:          +90 850 660 00 99
abuse-mailbox:  noc@vt.com.tr
fax-no:         +90 232 457 00 21
nic-hdl:        VT5050-RIPE
created:        2010-04-14T11:24:11Z
last-modified:  2015-01-07T21:13:34Z
source:         RIPE # Filtered
mnt-by:         MNT-VRGN

% Information related to '178.18.196.0/24AS43391'

route:          178.18.196.0/24
descr:          Vargonen Route
origin:         AS43391
mnt-by:         MNT-VRGN
created:        2014-12-01T23:24:24Z
last-modified:  2015-01-07T21:30:22Z
source:         RIPE # Filtered

% This query was served by the RIPE Database Query Service version 1.85.1 (DB-3)`

	//RSPL parse yazılabilir.
	//Mevcut bir RSPL parser kullanılabilir.
	//Daha kolay bir yol bulunabilir. bul.
}
