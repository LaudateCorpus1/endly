package cli

import (
	"github.com/viant/endly"
	"strings"
)

//GetPath returns hierarchical path to the latest Activity
func GetPath(activities *endly.Activities, runner *Runner, fullPath bool) (string, int) {
	var pathLength = 0
	var activityPath = make([]string, 0)

	for i, activity := range *activities {
		var tag = activity.FormatTag()
		pathLength += len(tag)
		serviceAction := ""
		if i+1 < len(*activities) || fullPath {
			serviceAction = runner.ColorText(activity.Service+"."+activity.Action, runner.ServiceActionColor)
			pathLength += len(activity.Service) + 1 + len(activity.Action)
		}

		tag = runner.ColorText(tag, runner.TagColor)
		if runner.InverseTag {
			tag = runner.ColorText(tag, "inverse")
		}
		activityPath = append(activityPath, runner.ColorText(activity.Workflow, runner.PathColor)+tag+serviceAction)
		pathLength += len(activity.Workflow)
	}

	var logPath = strings.Join(activityPath, runner.ColorText("|", "gray"))
	if len(*activities) > 0 {
		pathLength += (len(*activities) - 1)
	}
	return logPath, pathLength + 1
}
