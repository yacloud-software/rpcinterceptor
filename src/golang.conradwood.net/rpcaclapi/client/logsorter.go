package main

import (
	pb "golang.conradwood.net/apis/rpcaclapi"
)

type logsorter struct {
	Logs []*pb.LogEntry
}

func (ls *logsorter) ByRequestID(rid string) []*pb.LogEntry {
	var res []*pb.LogEntry
	for _, l := range ls.Logs {
		if l.Response.RequestID != rid {
			continue
		}
		res = append(res, l)
	}
	return res
}
func (ls *logsorter) RequestIDs() []*string {
	var res []*string
	for _, x := range ls.Logs {
		gotit := false
		for _, y := range res {
			if *y == x.Response.RequestID {
				gotit = true
				break
			}
		}
		if !gotit {
			res = append(res, &x.Response.RequestID)
		}
	}
	return res
}
func NewLogSorter(lel *pb.LogEntryList) *logsorter {
	l := &logsorter{}
	l.Logs = lel.Entries
	return l
}
