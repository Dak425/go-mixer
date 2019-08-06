package mixer

type IngestStreamAnalyticCollection []IngestStreamAnalytic
type IngestStreamAnalyticStringMap map[string]IngestStreamAnalytic
type IngestStreamAnalyticIntMap map[int]IngestStreamAnalytic
type IngestStreamAnalyticUIntMap map[uint]IngestStreamAnalytic

// A metric providing video details about a broadcast.
type IngestStreamAnalytic struct {
	AudioAvgNackDelay      uint `json:"audio_avg_nack_delay"`
	AudiooJitterBufferSize uint `json:"audioo_jitter_buffer_size"`
	AudiovAvgRecoveryDelay uint `json:"audiov_avg_recovery_delay"`
	AudiovBitrate          uint `json:"audiov_bitrate"`
	AudiovPacketRecovered  int  `json:"audiov_packet_recovered"`
	AudiovoPacketLoss      int  `json:"audiovo_packet_loss"`
	ChannelID              uint `json:"channel_id"`
	IngestID               uint `json:"ingest_id"`
	SubscribedDistNodes    uint `json:"subscribed_dist_nodes"`
	VideoAvgNackDelay      uint `json:"video_avg_nack_delay"`
	VideoAvgRecoveryDelay  uint `json:"video_avg_recovery_delay"`
	VideoBitrate           uint `json:"video_bitrate"`
	VideoJitterBufferSize  uint `json:"video_jitter_buffer_size"`
	VideoPacketLoss        int  `json:"video_packet_loss"`
	VideoPacketRecovered   int  `json:"video_packet_recovered"`
}
