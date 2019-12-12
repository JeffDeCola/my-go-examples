package main

const (
	genesisDataString = `
        {
        "txID": 0,
        "inputs": [
                {
                    "inID": 0,
                    "refTxID": -1,
                    "inPubKey": "",
                    "signature": "Hello World"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + foundersPubKey + `",
                    "value": 100000000
                }
            ]
        }`
	transaction1DataString = `
        {
        "txID": 1,
        "inputs": [
                {
                    "inID": 1,
                    "refTxID": 0,
                    "inPubKey": "` + foundersPubKey + `",
                    "signature": "` + foundersSig + `"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + jeffPubKey + `",
                    "value": 80000
                },
                {
                    "outPubKey": "` + foundersPubKey + `",
                    "value": 99920000
                }
            ]
        }`
	transaction2DataString = `
        {
        "txID": 2,
        "inputs": [
                {
                    "inID": 2,
                    "refTxID": 1,
                    "inPubKey": "` + jeffPubKey + `",
                    "signature": "` + jeffSig + `"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + mattPubKey + `",
                    "value": 50000
                },
                {
                    "outPubKey": "` + coinVaultPubKey + `",
                    "value": 500
                },
                {
                    "outPubKey": "` + jeffPubKey + `",
                    "value": 29500
                }
            ]
        }`
	transaction3DataString = `
        {
        "txID": 3,
        "inputs": [
                {
                    "inID": 3,
                    "refTxID": 1,
                    "inPubKey": "` + foundersPubKey + `",
                    "signature": "` + foundersSig + `"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + mattPubKey + `",
                    "value": 25000
                },
                {
                    "outPubKey": "` + foundersPubKey + `",
                    "value": 99670000
                }
            ]
        }`
	transaction4DataString = `
        {
        "txID": 4,
        "inputs": [
                {
                    "inID": 4,
                    "refTxID": 2,
                    "inPubKey": "` + mattPubKey + `",
                    "signature": "` + mattSig + `"
                },
                {
                    "inID": 5,
                    "refTxID": 3,
                    "inPubKey": "` + mattPubKey + `",
                    "signature": "` + mattSig + `"
                }
            ],
            "outputs": [
                {
                    "outPubKey": "` + jillPubKey + `",
                    "value": 2750000
                },
                {
                    "outPubKey": "` + mattPubKey + `",
                    "value": 25000
                }
            ]
        }`
)
