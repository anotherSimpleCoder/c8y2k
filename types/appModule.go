package types

import (
	"c8y2k/utils"
	"fmt"
	"os"
)

type AppModule struct {
	Imports []string
}

func (a AppModule) AddImport(moduleName string) {
	a.Imports = append(a.Imports, fmt.Sprintf("%sModule", utils.AngularString(moduleName)))
}

func (a AppModule) AppModuleContent() string {
	//Set header
	head := `
import {NgModule} from '@angular/core'
import {BrowserAnimationsModule} from '@angular/platform-browser/animations'
import { RouterModule as ngRouterModule } from '@angular/router'
import { CoreModule, BootstrapComponent, RouterModule } from '@c8y/ngx-components'
	`

	//Set NgModule imports
	ng_imports := "[BrowserAnimationModule,RouterModule.forRoot(),ngRouterModule.forRoot([], { enableTracing: false, useHash: true }),CoreModule.forRoot()"

	//Add modules
	for _, module := range a.Imports {
		head += fmt.Sprintf("import {%sModule} from './src/%s/%s.module'", utils.AngularString(module), module, module)
		ng_imports += fmt.Sprintf("%sModule,", utils.AngularString(module))
	}

	ng_imports += "]"

	return fmt.Sprintf(`
%s

@NgModule({
	imports: %s,
	bootstrap: [BootstrapComponent]
})
export class AppModule {}
	`, head, ng_imports)
}

func (a AppModule) Refresh() error {
	app_module_ts, err := os.Create("app.module.ts")
	if err != nil {
		return err
	}

	defer app_module_ts.Close()

	if _, err := app_module_ts.WriteString(a.AppModuleContent()); err != nil {
		return err
	}

	return nil
}
