package staticfile

// Code generated by go generate; DO NOT EDIT.

func init() {
	box.Add("query/AllAuditLogPerTimePeriod.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 65, 117, 100, 105, 116, 76, 111, 103, 76, 105, 115, 116, 81, 117, 101, 114, 121, 40, 36, 116, 105, 109, 101, 65, 103, 111, 58, 32, 68, 97, 116, 101, 84, 105, 109, 101, 41, 32, 123, 10, 9, 117, 115, 101, 114, 65, 117, 100, 105, 116, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 40, 102, 105, 108, 116, 101, 114, 115, 58, 32, 123, 32, 116, 105, 109, 101, 95, 103, 116, 58, 32, 36, 116, 105, 109, 101, 65, 103, 111, 32, 125, 41, 32, 123, 10, 9, 9, 101, 100, 103, 101, 115, 32, 123, 10, 9, 9, 9, 110, 111, 100, 101, 32, 123, 10, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 109, 101, 115, 115, 97, 103, 101, 10, 9, 9, 9, 9, 116, 105, 109, 101, 10, 9, 9, 9, 9, 115, 101, 118, 101, 114, 105, 116, 121, 10, 9, 9, 9, 9, 115, 116, 97, 116, 117, 115, 10, 9, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 32, 123, 10, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 125, 10, 9, 9, 125, 10, 9, 125, 10, 125, 10})
	box.Add("query/AllEventsPerTimePeriod.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 69, 118, 101, 110, 116, 83, 101, 114, 105, 101, 115, 76, 105, 115, 116, 81, 117, 101, 114, 121, 40, 36, 116, 105, 109, 101, 65, 103, 111, 58, 32, 68, 97, 116, 101, 84, 105, 109, 101, 41, 32, 123, 10, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 40, 102, 105, 108, 116, 101, 114, 115, 58, 32, 123, 32, 108, 97, 115, 116, 85, 112, 100, 97, 116, 101, 100, 95, 103, 116, 58, 32, 36, 116, 105, 109, 101, 65, 103, 111, 32, 125, 41, 32, 123, 10, 9, 9, 101, 100, 103, 101, 115, 32, 123, 10, 9, 9, 9, 110, 111, 100, 101, 32, 123, 10, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 102, 105, 100, 10, 9, 9, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 10, 9, 9, 9, 9, 108, 97, 115, 116, 85, 112, 100, 97, 116, 101, 100, 10, 9, 9, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 84, 121, 112, 101, 10, 9, 9, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 83, 116, 97, 116, 117, 115, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 73, 100, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 78, 97, 109, 101, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 84, 121, 112, 101, 10, 9, 9, 9, 9, 115, 101, 118, 101, 114, 105, 116, 121, 10, 9, 9, 9, 9, 112, 114, 111, 103, 114, 101, 115, 115, 10, 9, 9, 9, 9, 105, 115, 67, 97, 110, 99, 101, 108, 97, 98, 108, 101, 10, 9, 9, 9, 9, 105, 115, 80, 111, 108, 97, 114, 105, 115, 69, 118, 101, 110, 116, 83, 101, 114, 105, 101, 115, 10, 9, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 32, 123, 10, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 40, 102, 105, 114, 115, 116, 58, 32, 49, 41, 32, 123, 10, 9, 9, 9, 9, 9, 110, 111, 100, 101, 115, 32, 123, 10, 9, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 9, 109, 101, 115, 115, 97, 103, 101, 10, 9, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 125, 10, 9, 9, 125, 10, 9, 9, 112, 97, 103, 101, 73, 110, 102, 111, 32, 123, 10, 9, 9, 9, 101, 110, 100, 67, 117, 114, 115, 111, 114, 10, 9, 9, 9, 104, 97, 115, 78, 101, 120, 116, 80, 97, 103, 101, 10, 9, 9, 9, 104, 97, 115, 80, 114, 101, 118, 105, 111, 117, 115, 80, 97, 103, 101, 10, 9, 9, 125, 10, 9, 125, 10, 125, 10})
	box.Add("query/AllPolarisEventPerTimePeriod.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 65, 108, 108, 80, 111, 108, 97, 114, 105, 115, 69, 118, 101, 110, 116, 115, 80, 101, 114, 84, 105, 109, 101, 80, 101, 114, 105, 111, 100, 40, 36, 116, 105, 109, 101, 65, 103, 111, 58, 32, 68, 97, 116, 101, 84, 105, 109, 101, 44, 32, 36, 97, 102, 116, 101, 114, 58, 32, 83, 116, 114, 105, 110, 103, 41, 32, 123, 10, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 40, 10, 9, 9, 102, 105, 108, 116, 101, 114, 115, 58, 32, 123, 10, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 58, 32, 123, 32, 105, 100, 58, 32, 91, 34, 48, 48, 48, 48, 48, 48, 48, 48, 45, 48, 48, 48, 48, 45, 48, 48, 48, 48, 45, 48, 48, 48, 48, 45, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 93, 32, 125, 10, 9, 9, 9, 115, 116, 97, 114, 116, 84, 105, 109, 101, 95, 103, 116, 58, 32, 36, 116, 105, 109, 101, 65, 103, 111, 10, 9, 9, 9, 108, 97, 115, 116, 85, 112, 100, 97, 116, 101, 100, 95, 103, 116, 58, 32, 36, 116, 105, 109, 101, 65, 103, 111, 10, 9, 9, 125, 10, 9, 9, 102, 105, 114, 115, 116, 58, 32, 50, 48, 10, 9, 9, 97, 102, 116, 101, 114, 58, 32, 36, 97, 102, 116, 101, 114, 10, 9, 41, 32, 123, 10, 9, 9, 101, 100, 103, 101, 115, 32, 123, 10, 9, 9, 9, 110, 111, 100, 101, 32, 123, 10, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 102, 105, 100, 10, 9, 9, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 10, 9, 9, 9, 9, 108, 97, 115, 116, 85, 112, 100, 97, 116, 101, 100, 10, 9, 9, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 84, 121, 112, 101, 10, 9, 9, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 83, 116, 97, 116, 117, 115, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 73, 100, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 78, 97, 109, 101, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 84, 121, 112, 101, 10, 9, 9, 9, 9, 115, 101, 118, 101, 114, 105, 116, 121, 10, 9, 9, 9, 9, 112, 114, 111, 103, 114, 101, 115, 115, 10, 9, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 32, 123, 10, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 32, 123, 10, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 32, 123, 10, 9, 9, 9, 9, 9, 110, 111, 100, 101, 115, 32, 123, 10, 9, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 9, 109, 101, 115, 115, 97, 103, 101, 10, 9, 9, 9, 9, 9, 9, 116, 105, 109, 101, 10, 9, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 125, 10, 9, 9, 125, 10, 9, 9, 112, 97, 103, 101, 73, 110, 102, 111, 32, 123, 10, 9, 9, 9, 101, 110, 100, 67, 117, 114, 115, 111, 114, 10, 9, 9, 9, 104, 97, 115, 78, 101, 120, 116, 80, 97, 103, 101, 10, 9, 9, 9, 104, 97, 115, 80, 114, 101, 118, 105, 111, 117, 115, 80, 97, 103, 101, 10, 9, 9, 125, 10, 9, 125, 10, 125, 10})
	box.Add("query/EventDetails.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 69, 118, 101, 110, 116, 83, 101, 114, 105, 101, 115, 68, 101, 116, 97, 105, 108, 115, 81, 117, 101, 114, 121, 40, 36, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 58, 32, 85, 85, 73, 68, 33, 44, 32, 36, 99, 108, 117, 115, 116, 101, 114, 85, 117, 105, 100, 58, 32, 85, 85, 73, 68, 33, 41, 32, 123, 10, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 40, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 58, 32, 36, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 44, 32, 99, 108, 117, 115, 116, 101, 114, 85, 117, 105, 100, 58, 32, 36, 99, 108, 117, 115, 116, 101, 114, 85, 117, 105, 100, 41, 32, 123, 10, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 32, 123, 10, 9, 9, 9, 110, 111, 100, 101, 115, 32, 123, 10, 9, 9, 9, 9, 109, 101, 115, 115, 97, 103, 101, 10, 9, 9, 9, 9, 115, 116, 97, 116, 117, 115, 10, 9, 9, 9, 9, 116, 105, 109, 101, 10, 9, 9, 9, 9, 115, 101, 118, 101, 114, 105, 116, 121, 10, 9, 9, 9, 125, 10, 9, 9, 125, 10, 9, 9, 105, 100, 10, 9, 9, 102, 105, 100, 10, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 10, 9, 9, 111, 98, 106, 101, 99, 116, 73, 100, 10, 9, 9, 111, 98, 106, 101, 99, 116, 78, 97, 109, 101, 10, 9, 9, 111, 98, 106, 101, 99, 116, 84, 121, 112, 101, 10, 9, 9, 99, 108, 117, 115, 116, 101, 114, 32, 123, 10, 9, 9, 9, 105, 100, 10, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 125, 10, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 83, 116, 97, 116, 117, 115, 10, 9, 125, 10, 125, 10})
	box.Add("query/RadarEnabledClusters.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 82, 97, 100, 97, 114, 69, 110, 97, 98, 108, 101, 100, 67, 108, 117, 115, 116, 101, 114, 115, 32, 123, 10, 9, 114, 97, 100, 97, 114, 67, 108, 117, 115, 116, 101, 114, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 32, 123, 10, 9, 9, 110, 111, 100, 101, 115, 32, 123, 10, 9, 9, 9, 108, 97, 109, 98, 100, 97, 67, 111, 110, 102, 105, 103, 32, 123, 10, 9, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 73, 100, 10, 9, 9, 9, 9, 101, 110, 97, 98, 108, 101, 65, 117, 116, 111, 109, 97, 116, 105, 99, 70, 109, 100, 85, 112, 108, 111, 97, 100, 10, 9, 9, 9, 125, 10, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 125, 10, 9, 125, 10, 125, 10})
	box.Add("query/RadarEventsPerTimePeriod.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 82, 97, 100, 97, 114, 69, 118, 101, 110, 116, 80, 101, 114, 84, 105, 109, 101, 80, 101, 114, 105, 111, 100, 40, 36, 116, 105, 109, 101, 65, 103, 111, 58, 32, 68, 97, 116, 101, 84, 105, 109, 101, 41, 32, 123, 10, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 40, 102, 105, 108, 116, 101, 114, 115, 58, 32, 123, 32, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 84, 121, 112, 101, 58, 32, 65, 110, 111, 109, 97, 108, 121, 44, 32, 115, 116, 97, 114, 116, 84, 105, 109, 101, 95, 103, 116, 58, 32, 36, 116, 105, 109, 101, 65, 103, 111, 32, 125, 41, 32, 123, 10, 9, 9, 101, 100, 103, 101, 115, 32, 123, 10, 9, 9, 9, 110, 111, 100, 101, 32, 123, 10, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 102, 105, 100, 10, 9, 9, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 73, 100, 10, 9, 9, 9, 9, 108, 97, 115, 116, 85, 112, 100, 97, 116, 101, 100, 10, 9, 9, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 84, 121, 112, 101, 10, 9, 9, 9, 9, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 83, 116, 97, 116, 117, 115, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 73, 100, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 78, 97, 109, 101, 10, 9, 9, 9, 9, 111, 98, 106, 101, 99, 116, 84, 121, 112, 101, 10, 9, 9, 9, 9, 115, 101, 118, 101, 114, 105, 116, 121, 10, 9, 9, 9, 9, 112, 114, 111, 103, 114, 101, 115, 115, 10, 9, 9, 9, 9, 99, 108, 117, 115, 116, 101, 114, 32, 123, 10, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 110, 97, 109, 101, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 97, 99, 116, 105, 118, 105, 116, 121, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 32, 123, 10, 9, 9, 9, 9, 9, 110, 111, 100, 101, 115, 32, 123, 10, 9, 9, 9, 9, 9, 9, 105, 100, 10, 9, 9, 9, 9, 9, 9, 109, 101, 115, 115, 97, 103, 101, 10, 9, 9, 9, 9, 9, 9, 116, 105, 109, 101, 10, 9, 9, 9, 9, 9, 125, 10, 9, 9, 9, 9, 125, 10, 9, 9, 9, 125, 10, 9, 9, 125, 10, 9, 125, 10, 125, 10})
	box.Add("query/RadarTotalEventsPerTimePeriod.graphql", []byte{113, 117, 101, 114, 121, 32, 82, 98, 107, 76, 111, 103, 82, 97, 100, 97, 114, 69, 118, 101, 110, 116, 115, 80, 101, 114, 84, 105, 109, 101, 80, 101, 114, 105, 111, 100, 40, 36, 116, 105, 109, 101, 65, 103, 111, 58, 32, 68, 97, 116, 101, 84, 105, 109, 101, 41, 32, 123, 10, 9, 97, 99, 116, 105, 118, 105, 116, 121, 83, 101, 114, 105, 101, 115, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 40, 102, 105, 108, 116, 101, 114, 115, 58, 32, 123, 32, 108, 97, 115, 116, 65, 99, 116, 105, 118, 105, 116, 121, 84, 121, 112, 101, 58, 32, 65, 110, 111, 109, 97, 108, 121, 44, 32, 115, 116, 97, 114, 116, 84, 105, 109, 101, 95, 103, 116, 58, 32, 36, 116, 105, 109, 101, 65, 103, 111, 32, 125, 41, 32, 123, 10, 9, 9, 99, 111, 117, 110, 116, 10, 9, 125, 10, 125, 10})
}