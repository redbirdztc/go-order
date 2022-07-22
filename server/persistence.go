package main

// Persist persist message
func Persist(msg *message) error {
	return nil
}

// msgToBinary tansfer Message to binary data for persistance
// | total length | id | milli-time | content length | content | topic length | topic | message type |
// |------8B------|-8B-|-----8B-----|-------8B-------|-dynamic-|------8B------|dynamic|------8B------|
func msgToBinary(msg *message) (rs []byte, err error) {
	idBytes := int642bin(int64(msg.id))
	timeBytes := time2bin(&msg.ts)

	cttBytes := string2bin(msg.content)
	cttLenBytes := int642bin(int64(len(cttBytes)))

	topicBytes := string2bin(string(msg.topic))
	topicLenBytes := int642bin(int64(len(topicBytes)))

	tpBytes := int642bin(int64(msg.topicTp))

	totalLen := len(idBytes) + len(timeBytes) + len(cttLenBytes) +
		len(cttBytes) + len(topicLenBytes) + len(topicBytes) + len(tpBytes)
	totalLenBypte := int642bin(int64(totalLen))

	rs = make([]byte, 0, 8+totalLen)
	rs = append(rs, totalLenBypte...)
	rs = append(rs, idBytes...)
	rs = append(rs, timeBytes...)
	rs = append(rs, cttLenBytes...)
	rs = append(rs, cttBytes...)
	rs = append(rs, topicLenBytes...)
	rs = append(rs, topicBytes...)
	rs = append(rs, tpBytes...)

	return
}
