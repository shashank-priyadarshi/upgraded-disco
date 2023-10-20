package plugins

type Plugin struct{}

func(p *Plugin)Install(plugin string) error{return nil}
func(p *Plugin)Uninstall(plugin string) error{return nil}
func(p *Plugin)Upgrade(plugin string) error{return nil}
func(p *Plugin)Trigger(plugin string, payload interface{}) error{return nil}
func(p *Plugin)GetLastTrigger(plugin string, payload interface{}) interface{}{return nil}