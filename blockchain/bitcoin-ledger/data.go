// my-go-examples bitcoin-ledger data.go

package main

const (
	genesisTransactionString = `
        {
        "txID": 0,
        "inputs": [
                {
                    "inID": 0,
                    "refTxID": -1,
                    "inPubKey": "",
                    "signature": "Hello World, WElcome to the first transaction and block of jeffCoin"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + foundersPubKey + `",
                    "value": 100000000
                }
            ]
        }`
	txRequestMessageSignedDataString1 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 80000
                    }
                ]
            },
            "signature": "` + foundersSig + `"
        }`
	txRequestMessageSignedDataString2 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + jeffPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + mattPubKey + `",
                        "value": 50000
                    },
                    {
                        "destinationAddress": "` + coinVaultPubKey + `",
                        "value": 500
                    }
                ]
            },
            "signature": "` + jeffSig + `"
        }`
	txRequestMessageSignedDataString3 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + mattPubKey + `",
                        "value": 250000
                    }
                ]
            },
            "signature": "` + foundersSig + `"
        }`
	txRequestMessageSignedDataString4 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + mattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jillPubKey + `",
                        "value": 2750000
                    }
                ]
            },
            "signature": "` + mattSig + `"
        }`
	txRequestMessageSignedDataStringBad = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 80000
                    }
                ]
            },
            "signature": "Bad"
        }`
)
