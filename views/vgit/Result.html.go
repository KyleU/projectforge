// Code generated by qtc from "Result.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vgit/Result.html:1
package vgit

//line views/vgit/Result.html:1
import (
	"fmt"
	"strings"

	"projectforge.dev/projectforge/app"
	"projectforge.dev/projectforge/app/controller/cutil"
	"projectforge.dev/projectforge/app/lib/git"
	"projectforge.dev/projectforge/app/util"
	"projectforge.dev/projectforge/views/components"
	"projectforge.dev/projectforge/views/layout"
)

//line views/vgit/Result.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgit/Result.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgit/Result.html:13
type Result struct {
	layout.Basic
	Action *git.Action
	Result *git.Result
	URL    string
	Title  string
	Icon   string
}

//line views/vgit/Result.html:22
func (p *Result) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgit/Result.html:22
	qw422016.N().S(`
  <div class="card">
    <div class="right"><em>`)
//line views/vgit/Result.html:24
	qw422016.E().S(p.Result.Status)
//line views/vgit/Result.html:24
	qw422016.N().S(`</em></div>
    <h3>
      <a href="`)
//line views/vgit/Result.html:26
	qw422016.E().S(p.URL)
//line views/vgit/Result.html:26
	qw422016.N().S(`">`)
//line views/vgit/Result.html:26
	if p.Icon != "" {
//line views/vgit/Result.html:26
		components.StreamSVGIcon(qw422016, p.Icon, ps)
//line views/vgit/Result.html:26
		qw422016.N().S(` `)
//line views/vgit/Result.html:26
	}
//line views/vgit/Result.html:26
	qw422016.E().S(p.Title)
//line views/vgit/Result.html:26
	qw422016.N().S(`</a>: Git `)
//line views/vgit/Result.html:26
	qw422016.E().S(p.Action.Title)
//line views/vgit/Result.html:26
	qw422016.N().S(`
`)
//line views/vgit/Result.html:27
	if p.Result.DataString("branch") != "" {
//line views/vgit/Result.html:27
		qw422016.N().S(`      <em>(`)
//line views/vgit/Result.html:28
		qw422016.E().S(p.Result.DataString("branch"))
//line views/vgit/Result.html:28
		qw422016.N().S(`)</em>
`)
//line views/vgit/Result.html:29
	}
//line views/vgit/Result.html:29
	qw422016.N().S(`    </h3>
    <div class="mt">
      `)
//line views/vgit/Result.html:32
	streamstatusActions(qw422016, p.Result, false)
//line views/vgit/Result.html:32
	qw422016.N().S(`
    </div>
    `)
//line views/vgit/Result.html:34
	streamstatusDetail(qw422016, p.Result)
//line views/vgit/Result.html:34
	qw422016.N().S(`
  </div>
`)
//line views/vgit/Result.html:36
}

//line views/vgit/Result.html:36
func (p *Result) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgit/Result.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Result.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vgit/Result.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Result.html:36
}

//line views/vgit/Result.html:36
func (p *Result) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgit/Result.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Result.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vgit/Result.html:36
	qs422016 := string(qb422016.B)
//line views/vgit/Result.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Result.html:36
	return qs422016
//line views/vgit/Result.html:36
}

//line views/vgit/Result.html:38
func streamstatusActions(qw422016 *qt422016.Writer, r *git.Result, showProject bool) {
//line views/vgit/Result.html:38
	qw422016.N().S(`
`)
//line views/vgit/Result.html:39
	if showProject {
//line views/vgit/Result.html:39
		qw422016.N().S(`  <a href="/p/`)
//line views/vgit/Result.html:40
		qw422016.E().S(r.Project)
//line views/vgit/Result.html:40
		qw422016.N().S(`"><button>View Project</button></a>
`)
//line views/vgit/Result.html:41
	}
//line views/vgit/Result.html:42
	for _, t := range r.Actions() {
//line views/vgit/Result.html:42
		qw422016.N().S(`  <a href="/git/`)
//line views/vgit/Result.html:43
		qw422016.E().S(r.Project)
//line views/vgit/Result.html:43
		qw422016.N().S(`/`)
//line views/vgit/Result.html:43
		qw422016.E().S(t.Key)
//line views/vgit/Result.html:43
		qw422016.N().S(`" title="`)
//line views/vgit/Result.html:43
		qw422016.E().S(t.Description)
//line views/vgit/Result.html:43
		qw422016.N().S(`"><button>`)
//line views/vgit/Result.html:43
		qw422016.E().S(t.Title)
//line views/vgit/Result.html:43
		qw422016.N().S(`</button></a>
`)
//line views/vgit/Result.html:44
	}
//line views/vgit/Result.html:45
}

//line views/vgit/Result.html:45
func writestatusActions(qq422016 qtio422016.Writer, r *git.Result, showProject bool) {
//line views/vgit/Result.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Result.html:45
	streamstatusActions(qw422016, r, showProject)
//line views/vgit/Result.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Result.html:45
}

//line views/vgit/Result.html:45
func statusActions(r *git.Result, showProject bool) string {
//line views/vgit/Result.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Result.html:45
	writestatusActions(qb422016, r, showProject)
//line views/vgit/Result.html:45
	qs422016 := string(qb422016.B)
//line views/vgit/Result.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Result.html:45
	return qs422016
//line views/vgit/Result.html:45
}

//line views/vgit/Result.html:47
func streamstatusDetail(qw422016 *qt422016.Writer, r *git.Result) {
//line views/vgit/Result.html:47
	qw422016.N().S(`
  <div class="overflow full-width">
    <table class="mt min-200">
      <tbody>
        <tr>
          <th class="shrink">Status</th>
          <td>`)
//line views/vgit/Result.html:53
	qw422016.E().S(r.Status)
//line views/vgit/Result.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink">Branch</th>
          <td>`)
//line views/vgit/Result.html:57
	qw422016.E().S(r.DataString("branch"))
//line views/vgit/Result.html:57
	qw422016.N().S(`</td>
        </tr>
`)
//line views/vgit/Result.html:59
	if r.DataString("commitMessage") != "" {
//line views/vgit/Result.html:59
		qw422016.N().S(`        <tr>
          <th class="shrink">Commit Message</th>
          <td>`)
//line views/vgit/Result.html:62
		qw422016.E().S(r.DataString("commitMessage"))
//line views/vgit/Result.html:62
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vgit/Result.html:64
	}
//line views/vgit/Result.html:65
	if r.DataString("commit") != "" {
//line views/vgit/Result.html:65
		qw422016.N().S(`        <tr>
          <th class="shrink">Commit Results</th>
          <td><pre>`)
//line views/vgit/Result.html:68
		qw422016.N().S(r.DataString("commit"))
//line views/vgit/Result.html:68
		qw422016.N().S(`</pre></td>
        </tr>
`)
//line views/vgit/Result.html:70
	}
//line views/vgit/Result.html:71
	if delta := r.DataInt("commitsAhead"); delta > 0 {
//line views/vgit/Result.html:71
		qw422016.N().S(`        <tr>
          <th class="shrink">Ahead By</th>
          <td>
            <div class="right"><a href="/git/`)
//line views/vgit/Result.html:75
		qw422016.E().S(r.Project)
//line views/vgit/Result.html:75
		qw422016.N().S(`/push"><button>Push</button></a></div>
            `)
//line views/vgit/Result.html:76
		qw422016.E().S(util.StringPlural(delta, "commit"))
//line views/vgit/Result.html:76
		qw422016.N().S(`
          </td>
        </tr>
`)
//line views/vgit/Result.html:79
	}
//line views/vgit/Result.html:80
	if delta := r.DataInt("commitsBehind"); delta > 0 {
//line views/vgit/Result.html:80
		qw422016.N().S(`        <tr>
          <th class="shrink">Behind By</th>
          <td>
            <div class="right"><a href="/git/`)
//line views/vgit/Result.html:84
		qw422016.E().S(r.Project)
//line views/vgit/Result.html:84
		qw422016.N().S(`/pull"><button>Pull</button></a></div>
            `)
//line views/vgit/Result.html:85
		qw422016.E().S(util.StringPlural(delta, "commit"))
//line views/vgit/Result.html:85
		qw422016.N().S(`
          </td>
        </tr>
`)
//line views/vgit/Result.html:88
	}
//line views/vgit/Result.html:89
	if len(r.DataStringArray("logs")) > 0 {
//line views/vgit/Result.html:89
		qw422016.N().S(`        <tr>
          <th class="shrink">Logs</th>
          <td>
            <div class="overflow full-width">
              <table>
                <tbody>
`)
//line views/vgit/Result.html:96
		for idx, l := range r.DataStringArray("logs") {
//line views/vgit/Result.html:96
			qw422016.N().S(`                  <tr><td class="shrink"><em>`)
//line views/vgit/Result.html:97
			qw422016.N().D(idx + 1)
//line views/vgit/Result.html:97
			qw422016.N().S(`</em>:</td><td>`)
//line views/vgit/Result.html:97
			qw422016.E().S(l)
//line views/vgit/Result.html:97
			qw422016.N().S(`</td></tr>
`)
//line views/vgit/Result.html:98
		}
//line views/vgit/Result.html:98
		qw422016.N().S(`                </tbody>
              </table>
            </div>
          </td>
        </tr>
`)
//line views/vgit/Result.html:104
	}
//line views/vgit/Result.html:104
	qw422016.N().S(`
        <tr>
          <th class="shrink">Dirty Files</th>
          <td>
`)
//line views/vgit/Result.html:109
	dirt := r.DataStringArray("dirty")

//line views/vgit/Result.html:110
	if len(dirt) > 0 {
//line views/vgit/Result.html:110
		qw422016.N().S(`            <div class="right"><a href="/git/`)
//line views/vgit/Result.html:111
		qw422016.E().S(r.Project)
//line views/vgit/Result.html:111
		qw422016.N().S(`/commit"><button>Commit</button></a></div>
            <ul>
`)
//line views/vgit/Result.html:113
		for _, d := range dirt {
//line views/vgit/Result.html:113
			qw422016.N().S(`              <li>`)
//line views/vgit/Result.html:114
			qw422016.E().S(d)
//line views/vgit/Result.html:114
			qw422016.N().S(`</li>
`)
//line views/vgit/Result.html:115
		}
//line views/vgit/Result.html:115
		qw422016.N().S(`            </ul>
`)
//line views/vgit/Result.html:117
	} else {
//line views/vgit/Result.html:117
		qw422016.N().S(`            <em>none</em>
`)
//line views/vgit/Result.html:119
	}
//line views/vgit/Result.html:119
	qw422016.N().S(`          </td>
        </tr>

`)
//line views/vgit/Result.html:123
	if hist := r.History(); hist != nil {
//line views/vgit/Result.html:123
		qw422016.N().S(`        `)
//line views/vgit/Result.html:124
		streamdisplayHistory(qw422016, r.Project, hist)
//line views/vgit/Result.html:124
		qw422016.N().S(`
`)
//line views/vgit/Result.html:125
	} else if cleanData := r.CleanData(); len(cleanData) > 0 {
//line views/vgit/Result.html:125
		qw422016.N().S(`        <tr>
          <th class="shrink">Data</th>
          <td>`)
//line views/vgit/Result.html:128
		components.StreamJSON(qw422016, cleanData)
//line views/vgit/Result.html:128
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vgit/Result.html:130
	}
//line views/vgit/Result.html:130
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vgit/Result.html:134
}

//line views/vgit/Result.html:134
func writestatusDetail(qq422016 qtio422016.Writer, r *git.Result) {
//line views/vgit/Result.html:134
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Result.html:134
	streamstatusDetail(qw422016, r)
//line views/vgit/Result.html:134
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Result.html:134
}

//line views/vgit/Result.html:134
func statusDetail(r *git.Result) string {
//line views/vgit/Result.html:134
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Result.html:134
	writestatusDetail(qb422016, r)
//line views/vgit/Result.html:134
	qs422016 := string(qb422016.B)
//line views/vgit/Result.html:134
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Result.html:134
	return qs422016
//line views/vgit/Result.html:134
}

//line views/vgit/Result.html:136
func streamdisplayHistory(qw422016 *qt422016.Writer, key string, res *git.HistoryResult) {
//line views/vgit/Result.html:136
	qw422016.N().S(`
`)
//line views/vgit/Result.html:137
	if res.Args.Since != nil {
//line views/vgit/Result.html:137
		qw422016.N().S(`  <tr>
    <th>Since</th>
    <td>`)
//line views/vgit/Result.html:140
		qw422016.E().S(util.TimeToFull(res.Args.Since))
//line views/vgit/Result.html:140
		qw422016.N().S(`</td>
  </tr>
`)
//line views/vgit/Result.html:142
	}
//line views/vgit/Result.html:143
	if len(res.Args.Authors) > 0 {
//line views/vgit/Result.html:143
		qw422016.N().S(`  <tr>
    <th>Authors</th>
    <td>`)
//line views/vgit/Result.html:146
		qw422016.E().S(strings.Join(res.Args.Authors, ", "))
//line views/vgit/Result.html:146
		qw422016.N().S(`</td>
  </tr>
`)
//line views/vgit/Result.html:148
	}
//line views/vgit/Result.html:149
	if res.Args.Limit > 0 {
//line views/vgit/Result.html:149
		qw422016.N().S(`  <tr>
    <th>Limit</th>
    <td>`)
//line views/vgit/Result.html:152
		qw422016.N().D(res.Args.Limit)
//line views/vgit/Result.html:152
		qw422016.N().S(`</td>
  </tr>
`)
//line views/vgit/Result.html:154
	}
//line views/vgit/Result.html:155
	if len(res.Entries) > 0 {
//line views/vgit/Result.html:155
		qw422016.N().S(`  `)
//line views/vgit/Result.html:156
		streamrenderPlot(qw422016, key, res.Entries)
//line views/vgit/Result.html:156
		qw422016.N().S(`
  <tr>
    <th>Entries (`)
//line views/vgit/Result.html:158
		qw422016.N().D(len(res.Entries))
//line views/vgit/Result.html:158
		qw422016.N().S(`)</th>
    <td>
`)
//line views/vgit/Result.html:160
		for _, e := range res.Entries {
//line views/vgit/Result.html:160
			qw422016.N().S(`        `)
//line views/vgit/Result.html:161
			streamdisplayEntry(qw422016, res.Args, e)
//line views/vgit/Result.html:161
			qw422016.N().S(`
`)
//line views/vgit/Result.html:162
		}
//line views/vgit/Result.html:162
		qw422016.N().S(`    </td>
  </tr>
`)
//line views/vgit/Result.html:165
	}
//line views/vgit/Result.html:166
}

//line views/vgit/Result.html:166
func writedisplayHistory(qq422016 qtio422016.Writer, key string, res *git.HistoryResult) {
//line views/vgit/Result.html:166
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Result.html:166
	streamdisplayHistory(qw422016, key, res)
//line views/vgit/Result.html:166
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Result.html:166
}

//line views/vgit/Result.html:166
func displayHistory(key string, res *git.HistoryResult) string {
//line views/vgit/Result.html:166
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Result.html:166
	writedisplayHistory(qb422016, key, res)
//line views/vgit/Result.html:166
	qs422016 := string(qb422016.B)
//line views/vgit/Result.html:166
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Result.html:166
	return qs422016
//line views/vgit/Result.html:166
}

//line views/vgit/Result.html:168
func streamrenderPlot(qw422016 *qt422016.Writer, key string, e git.HistoryEntries) {
//line views/vgit/Result.html:168
	qw422016.N().S(`
  `)
//line views/vgit/Result.html:169
	components.StreamPlotAssets(qw422016)
//line views/vgit/Result.html:169
	qw422016.N().S(`
  <tr>
    <th>Commit Summary</th>
    <td>
      <div id="`)
//line views/vgit/Result.html:173
	qw422016.E().S(key)
//line views/vgit/Result.html:173
	qw422016.N().S(`-plot"></div>
      <div class="mt">
`)
//line views/vgit/Result.html:175
	authors := e.Authors()

//line views/vgit/Result.html:175
	qw422016.N().S(`        <table>
          <thead>
            <tr>
              <th>Email</th>
              <th>Name</th>
              <th>Commits</th>
            </tr>
          </thead>
          <tbody>
`)
//line views/vgit/Result.html:185
	for _, x := range authors {
//line views/vgit/Result.html:185
		qw422016.N().S(`            <tr>
              <td class="shrink">`)
//line views/vgit/Result.html:187
		qw422016.E().S(x.Key)
//line views/vgit/Result.html:187
		qw422016.N().S(`</td>
              <td>`)
//line views/vgit/Result.html:188
		qw422016.E().S(x.Name)
//line views/vgit/Result.html:188
		qw422016.N().S(`</td>
              <td>`)
//line views/vgit/Result.html:189
		qw422016.N().D(x.Count)
//line views/vgit/Result.html:189
		qw422016.N().S(`</td>
            </tr>
`)
//line views/vgit/Result.html:191
	}
//line views/vgit/Result.html:191
	qw422016.N().S(`          </tbody>
        </table>
      </div>
      <script type="module">
        const authors = `)
//line views/vgit/Result.html:196
	qw422016.N().S(util.ToJSON(authors))
//line views/vgit/Result.html:196
	qw422016.N().S(`;

        function createPlot(div) {
          function render() {
            const height = `)
//line views/vgit/Result.html:200
	qw422016.N().D((len(authors) * 24) + 48)
//line views/vgit/Result.html:200
	qw422016.N().S(`;
            const width = div.clientWidth;
            const plot = Plot.plot({
              height,
              width,
              color: {
                legend: false,
              },
              marks: [
                Plot.barX(authors, {
                  x: "count",
                  fill: "key",
                  y: "key",
                  title: d => `)
//line views/vgit/Result.html:200
	qw422016.N().S("`")
//line views/vgit/Result.html:200
	qw422016.N().S(`${d.name}: ${d.count}`)
//line views/vgit/Result.html:200
	qw422016.N().S("`")
//line views/vgit/Result.html:200
	qw422016.N().S(`,
                  marginLeft: 256,
                })
              ]
            });
            div.innerHTML = "";
            div.appendChild(plot);
          }

          window.addEventListener('resize', render);
          render();

          return () => {
            window.removeEventListener('resize', render);
          };
        }

        function go() {
          const div = document.querySelector("#`)
//line views/vgit/Result.html:231
	qw422016.E().S(key)
//line views/vgit/Result.html:231
	qw422016.N().S(`-plot");
          const cancel = createPlot(div);
        }

        document.addEventListener("DOMContentLoaded", go);
      </script>
    </td>
  </tr>
`)
//line views/vgit/Result.html:239
}

//line views/vgit/Result.html:239
func writerenderPlot(qq422016 qtio422016.Writer, key string, e git.HistoryEntries) {
//line views/vgit/Result.html:239
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Result.html:239
	streamrenderPlot(qw422016, key, e)
//line views/vgit/Result.html:239
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Result.html:239
}

//line views/vgit/Result.html:239
func renderPlot(key string, e git.HistoryEntries) string {
//line views/vgit/Result.html:239
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Result.html:239
	writerenderPlot(qb422016, key, e)
//line views/vgit/Result.html:239
	qs422016 := string(qb422016.B)
//line views/vgit/Result.html:239
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Result.html:239
	return qs422016
//line views/vgit/Result.html:239
}

//line views/vgit/Result.html:241
func streamdisplayEntry(qw422016 *qt422016.Writer, args *git.HistoryArgs, e *git.HistoryEntry) {
//line views/vgit/Result.html:241
	qw422016.N().S(`
`)
//line views/vgit/Result.html:243
	u := "?path=" + args.Path
	if args.Since != nil {
		u += "&since=" + util.TimeToYMD(args.Since)
	}
	if len(args.Authors) > 0 {
		u += "&authors=" + strings.Join(args.Authors, ",")
	}
	if args.Limit > 0 {
		u += "&limit=" + fmt.Sprint(args.Limit)
	}
	u += "&commit=" + e.SHA

//line views/vgit/Result.html:254
	qw422016.N().S(`  <div class="padding border mb">
    <div class="right"><em>`)
//line views/vgit/Result.html:256
	qw422016.E().S(util.TimeToFull(&e.Occurred))
//line views/vgit/Result.html:256
	qw422016.N().S(`</em></div>
    <h4><a href="`)
//line views/vgit/Result.html:257
	qw422016.E().S(u)
//line views/vgit/Result.html:257
	qw422016.N().S(`">`)
//line views/vgit/Result.html:257
	qw422016.E().S(e.SHA)
//line views/vgit/Result.html:257
	qw422016.N().S(`</a></h4>
    <div class="mt">`)
//line views/vgit/Result.html:258
	qw422016.E().S(e.Message)
//line views/vgit/Result.html:258
	qw422016.N().S(`</div>
    <div class="mt">
      <em>
        -
        `)
//line views/vgit/Result.html:262
	if e.AuthorName != "" {
//line views/vgit/Result.html:262
		qw422016.E().S(e.AuthorName)
//line views/vgit/Result.html:262
	}
//line views/vgit/Result.html:262
	qw422016.N().S(`
        `)
//line views/vgit/Result.html:263
	if e.AuthorEmail == "" {
//line views/vgit/Result.html:263
		qw422016.N().S(`unknown author`)
//line views/vgit/Result.html:263
	} else {
//line views/vgit/Result.html:263
		qw422016.N().S(`(<a href="mailto:`)
//line views/vgit/Result.html:263
		qw422016.E().S(e.AuthorEmail)
//line views/vgit/Result.html:263
		qw422016.N().S(`">`)
//line views/vgit/Result.html:263
		qw422016.E().S(e.AuthorEmail)
//line views/vgit/Result.html:263
		qw422016.N().S(`</a>)`)
//line views/vgit/Result.html:263
	}
//line views/vgit/Result.html:263
	qw422016.N().S(`
      </em>
    </div>
`)
//line views/vgit/Result.html:266
	if len(e.Files) > 0 {
//line views/vgit/Result.html:266
		qw422016.N().S(`    <div class="mt">
      <ul>
`)
//line views/vgit/Result.html:269
		for _, f := range e.Files {
//line views/vgit/Result.html:269
			qw422016.N().S(`        <li>`)
//line views/vgit/Result.html:270
			qw422016.E().S(f.File)
//line views/vgit/Result.html:270
			qw422016.N().S(`</li>
`)
//line views/vgit/Result.html:271
		}
//line views/vgit/Result.html:271
		qw422016.N().S(`      </ul>
    </div>
`)
//line views/vgit/Result.html:274
	}
//line views/vgit/Result.html:274
	qw422016.N().S(`  </div>
`)
//line views/vgit/Result.html:276
}

//line views/vgit/Result.html:276
func writedisplayEntry(qq422016 qtio422016.Writer, args *git.HistoryArgs, e *git.HistoryEntry) {
//line views/vgit/Result.html:276
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgit/Result.html:276
	streamdisplayEntry(qw422016, args, e)
//line views/vgit/Result.html:276
	qt422016.ReleaseWriter(qw422016)
//line views/vgit/Result.html:276
}

//line views/vgit/Result.html:276
func displayEntry(args *git.HistoryArgs, e *git.HistoryEntry) string {
//line views/vgit/Result.html:276
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgit/Result.html:276
	writedisplayEntry(qb422016, args, e)
//line views/vgit/Result.html:276
	qs422016 := string(qb422016.B)
//line views/vgit/Result.html:276
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgit/Result.html:276
	return qs422016
//line views/vgit/Result.html:276
}
