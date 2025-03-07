package action

import (
	"context"
	"fmt"

	"github.com/robert-nix/ansihtml"

	"projectforge.dev/projectforge/app/lib/websocket"
	"projectforge.dev/projectforge/app/project"
	"projectforge.dev/projectforge/app/util"
)

func onStart(_ context.Context, pm *PrjAndMods, ret *Result) *Result {
	if ret == nil {
		ret = newResult(TypeDebug, pm.Prj, pm.Cfg, pm.Logger)
	}

	cmd := fmt.Sprintf("%s%s -v server source:%s", project.DebugOutputDir, pm.Prj.Executable(), util.AppKey)
	exec := pm.XSvc.NewExec(pm.Prj.Key, cmd, pm.Prj.Path, false, pm.Prj.Info.EnvVars...)
	exec.Link = pm.Prj.WebPath()
	w := func(key string, b []byte) error {
		m := util.ValueMap{"msg": string(b), "html": string(ansihtml.ConvertToHTML(b))}
		msg := &websocket.Message{Channel: key, Cmd: "output", Param: util.ToJSONBytes(m, true)}
		if pm.SSvc == nil {
			return nil
		}
		return pm.SSvc.WriteChannel(msg, pm.Logger)
	}
	err := exec.Start(w)
	if err != nil {
		return errorResult(err, TypeDebug, pm.Cfg, pm.Logger)
	}
	ret.Data = fmt.Sprintf("/admin/exec/%s/%d", pm.Prj.Key, 1)

	return ret
}
