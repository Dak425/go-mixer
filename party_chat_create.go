package structs

type PartyChatCreateCollection []PartyChatCreate
type PartyChatCreateStringMap map[string]PartyChatCreate
type PartyChatCreateIntMap map[int]PartyChatCreate
type PartyChatCreateUIntMap map[uint]PartyChatCreate

// Sent in the request to create a new party chat channel.
type PartyChatCreate struct {
	DeviceUUID                      string `json:"deviceUUID"`
	ProtocolVersion                 int    `json:"protocolVersion"`
	WebRtcDtlsCertificateAlgorithm  string `json:"webRtcDtlsCertificateAlgorithm"`
	WebRtcDtlsCertificateThumbprint string `json:"webRtcDtlsCertificateThumbprint"`
	WebRtcIcePwd                    string `json:"webRtcIcePwd"`
	WebRtcIceUfrag                  string `json:"webRtcIceUfrag"`
}
