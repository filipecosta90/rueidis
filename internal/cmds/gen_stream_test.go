// Code generated by go generate; DO NOT EDIT

package cmds

import "testing"

func stream0(s Builder) {
	s.Xack().Key("1").Group("1").Id("1").Id("1").Build()
	s.Xadd().Key("1").Nomkstream().Maxlen().Exact().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Maxlen().Exact().Threshold("1").Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Maxlen().Almost().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Maxlen().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Minid().Exact().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Minid().Almost().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Minid().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Nomkstream().Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Maxlen().Exact().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Minid().Exact().Threshold("1").Limit(1).Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xadd().Key("1").Id("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Xautoclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Start("1").Count(1).Justid().Build()
	s.Xautoclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Start("1").Count(1).Build()
	s.Xautoclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Start("1").Justid().Build()
	s.Xautoclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Start("1").Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Force().Justid().Lastid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Force().Justid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Force().Lastid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Force().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Justid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Lastid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Retrycount(1).Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Force().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Justid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Lastid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Time(1).Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Retrycount(1).Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Force().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Justid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Lastid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Idle(1).Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Time(1).Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Retrycount(1).Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Force().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Justid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Lastid().Build()
	s.Xclaim().Key("1").Group("1").Consumer("1").MinIdleTime("1").Id("1").Id("1").Build()
	s.Xdel().Key("1").Id("1").Id("1").Build()
	s.XgroupCreate().Key("1").Group("1").Id("1").Mkstream().Entriesread(1).Build()
	s.XgroupCreate().Key("1").Group("1").Id("1").Mkstream().Build()
	s.XgroupCreate().Key("1").Group("1").Id("1").Entriesread(1).Build()
	s.XgroupCreate().Key("1").Group("1").Id("1").Build()
	s.XgroupCreateconsumer().Key("1").Group("1").Consumer("1").Build()
	s.XgroupDelconsumer().Key("1").Group("1").Consumername("1").Build()
	s.XgroupDestroy().Key("1").Group("1").Build()
	s.XgroupHelp().Build()
	s.XgroupSetid().Key("1").Group("1").Id("1").Entriesread(1).Build()
	s.XgroupSetid().Key("1").Group("1").Id("1").Build()
	s.XinfoConsumers().Key("1").Group("1").Build()
	s.XinfoGroups().Key("1").Build()
	s.XinfoHelp().Build()
	s.XinfoStream().Key("1").Full().Count(1).Build()
	s.XinfoStream().Key("1").Full().Build()
	s.XinfoStream().Key("1").Build()
	s.Xlen().Key("1").Build()
	s.Xpending().Key("1").Group("1").Idle(1).Start("1").End("1").Count(1).Consumer("1").Build()
	s.Xpending().Key("1").Group("1").Idle(1).Start("1").End("1").Count(1).Build()
	s.Xpending().Key("1").Group("1").Start("1").End("1").Count(1).Build()
	s.Xpending().Key("1").Group("1").Build()
	s.Xrange().Key("1").Start("1").End("1").Count(1).Build()
	s.Xrange().Key("1").Start("1").End("1").Build()
	s.Xread().Count(1).Block(1).Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xread().Count(1).Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xread().Block(1).Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xread().Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Count(1).Block(1).Noack().Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Count(1).Block(1).Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Count(1).Noack().Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Count(1).Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Block(1).Noack().Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Noack().Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xreadgroup().Group("1", "1").Streams().Key("1").Key("1").Id("1").Id("1").Build()
	s.Xrevrange().Key("1").End("1").Start("1").Count(1).Build()
	s.Xrevrange().Key("1").End("1").Start("1").Build()
	s.Xsetid().Key("1").LastId("1").Entriesadded(1).Maxdeletedid("1").Build()
	s.Xsetid().Key("1").LastId("1").Entriesadded(1).Build()
	s.Xsetid().Key("1").LastId("1").Maxdeletedid("1").Build()
	s.Xsetid().Key("1").LastId("1").Build()
	s.Xtrim().Key("1").Maxlen().Exact().Threshold("1").Limit(1).Build()
	s.Xtrim().Key("1").Maxlen().Exact().Threshold("1").Build()
	s.Xtrim().Key("1").Maxlen().Almost().Threshold("1").Build()
	s.Xtrim().Key("1").Maxlen().Threshold("1").Build()
	s.Xtrim().Key("1").Minid().Exact().Threshold("1").Build()
	s.Xtrim().Key("1").Minid().Almost().Threshold("1").Build()
	s.Xtrim().Key("1").Minid().Threshold("1").Build()
}

func TestCommand_InitSlot_stream(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { stream0(s) })
}

func TestCommand_NoSlot_stream(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { stream0(s) })
}
