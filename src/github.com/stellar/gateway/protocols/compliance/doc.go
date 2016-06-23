/*
Package compliance contains message structures required to implement compliance protocol.

You can check compliance protocol doc here: https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html

Sending a payment

When you want to send payment using compliance protocol you need to send additional `sender` and `extra_memo` fields
to the `/payment` endpoint.

    Stellar    Bridge                            Compliance                             Acme Bank
       |          |                                    |                                    |
       |          |      compliance.HandlerSend()      |                                    |
       |          |               /send                |                                    |
       |          |       compliance.SendRequest       |                                    |
       |          □----------------------------------->□~~~ - Builds transaction            |
       |          □                                    □    - Gets AuthServer from          |
       |          □                                    □      stellar.toml file             |
       |          □                                    □    - Sends request to receiver     |
       |          □                                    □      AuthServer                    |
       |          □                                    □                                    |
       |          □                                    □       compliance.AuthRequest       □
       |          □                                    □----------------------------------->□~~~ - Performs KYC
       |          □                                    □                                    □    - Saves the memo preimage
       |          □                                    □       compliance.AuthResponse      □
       |          □      compliance.SendResponse       □<-----------------------------------□
       |          □<-----------------------------------□                                    |
       |          □~~~ - Signs the transaction         |                                    |
       |          □      returned by compliance        |                                    |
       |          □      server                        |                                    |
       |          □    - Submits a transaction to      |                                    |
       |          □      the Stellar network           |                                    |
       |<---------□                                    |                                    |
       |          |                                    |                                    |

Receiving a payment

When you want to use compliance protocol when receiving payments...

    Stellar    Bridge                            Compliance                             Acme Bank
       |          |                                    |                                    |
       |          |                                    |        compliance.AuthRequest      □
       |          |                                    □<-----------------------------------□~~~ - Other organization sends
       |          |                                    □~~~ - You perform KYC               □      request to your AuthServer
       |          |                                    □    - You save the memo preimage    □
       |          |                                    □                                    □
       |          |                                    □      compliance.AuthResponse       □
       |          |                                    □----------------------------------->□
       |          |                                    |                                    |
       |--------->□~~~ - If you respond with           |                                    |
       |          □      AUTH_STATUS_OK to AuthRequest |                                    |
       |          □      transaction will be sent      |                                    |
       |          □    - Payment is received           |                                    |
       |          □      by payment listener           |                                    |
       |          □                                    |                                    |
       |          □    compliance.HandlerReceive()     |                                    |
       |          □            /receive                |                                    |
       |          □     compliance.ReceiveRequest      |                                    |
       |          □----------------------------------->□~~~ - Compliance server will load   |
       |          □                                    □      memo preimage that was saved  |
       |          □                                    □      earlier                       |
       |          □                                    □                                    |
       |          □    compliance.ReceiveResponse      □                                    |
       |          □<-----------------------------------□                                    |
       |          □                                    |                                    |
       |          □~~~ - Bridge server will send       |                                    |
       |          □      a receive callback with       |                                    |
       |          □      `extra_memo` preimage         |                                    |
       |          □                                    |                                    |
       |          |                                    |                                    |

*/
package compliance
