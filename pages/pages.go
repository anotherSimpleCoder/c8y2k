package pages

import (
	"c8y2k/utils"
	"fmt"
	"os"
	"os/exec"
)

func Help() string {

	return `
c8y2k - A little tool for making Cumulocity apps easier

	Syntax:
		c8y2k <command>
	
	Command:
		help			Opens this page
		new			Creates a new Cumulocity Web SDK project
		new-component		Creates a new Component
		new-widget		Creates a new widget component
`
}

func NewProject() string {
	var projName string
	fmt.Print("Enter your project name: ")
	if _, err := fmt.Scanln(&projName); err != nil {
		return err.Error()
	}

	fmt.Println("Starting c8y Web SDK....")

	cmd := exec.Command("npx", "@c8y/cli@latest", "new", projName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err.Error()
	}

	if err := cmd.Wait(); err != nil {
		return err.Error()
	}

	if err := os.Chdir(projName); err != nil {
		return err.Error()
	}

	fmt.Println("Running npm install...")
	cmd = exec.Command("npm", "install")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err.Error()
	}

	if err := cmd.Wait(); err != nil {
		return err.Error()
	}

	return "Project successfully created"
}

func NewComponent() string {
	var compName string
	fmt.Print("Enter your component name: ")
	if _, err := fmt.Scanln(&compName); err != nil {
		return err.Error()
	}

	os.Mkdir("src", 0755)

	//Create component folder
	if err := os.Mkdir("src/"+compName, 0755); err != nil {
		return err.Error()
	}

	//Create template html
	if template_html_file, err := os.Create(fmt.Sprintf("src/%s/%s.component.html", compName, compName)); err != nil {
		return err.Error()
	} else {
		defer template_html_file.Close()

		template_html_file.WriteString(fmt.Sprintf("<c8y-title>{{'%s'}}</c8y-title>", utils.AngularString(compName)))
	}

	//Create style css
	if _, err := os.Create(fmt.Sprintf("src/%s/%s.component.css", compName, compName)); err != nil {
		return err.Error()
	}

	//Create component ts
	if comp_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.component.ts", compName, compName)); err != nil {
		return err.Error()
	} else {
		defer comp_ts_file.Close()

		comp_ts_file.WriteString(fmt.Sprintf(`
import {Component, OnInit} from '@angular/core'
	
@Component({
	selector: '%s',
	templateUrl: '%s.component.html',
	styleUrls: ['%s.component.css']
})
export class %sComponent implements OnInit {
	contructor(){}
	
	ngOnInit() {}
}
		`, compName, compName, compName, utils.AngularString(compName)))
	}

	//Create factory
	if factory_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.factory.ts", compName, compName)); err != nil {
		return err.Error()
	} else {
		defer factory_ts_file.Close()

		factory_ts_file.WriteString(fmt.Sprintf(`
import {Injectable} from '@angular/core'
import {NavigatorNode, NavigatorNodeFactory} from '@c8y/ngx-components'

@Injectable()
export class %sNavigationFactory implements NavigatorNodeFactory {
	get() {
		return new NavigatorNode({
			label: '%s',
			icon: 'robot',
			path: '%s',
			prioirty: 100
		})
	}
}
		`, utils.AngularString(compName), utils.AngularString(compName), compName))
	}

	//Create model
	if _, err := os.Create(fmt.Sprintf("src/%s/%s.model.ts", compName, compName)); err != nil {
		return err.Error()
	}

	//Create module
	if module_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.module.ts", compName, compName)); err != nil {
		return err.Error()
	} else {
		defer module_ts_file.Close()

		module_ts_file.WriteString(fmt.Sprintf(`
import {NgModule} from '@angular/core'
import {RouterModule,Routes} from '@angular/router'
import {CoreModule,hookNavigator} from '@c8y/ngx-components'

import {%sComponent} from './%s.component'
import {%sNavigationFactory} from './%s.factory'

const routes: Routes = [
	{
		path: '',
		pathMatch: 'full'
	},

	{
		path: '%s',
		component: %sComponent
	}
]
		
@NgModule({
	imports: [RouterModule.forChild(routes), CoreModule],
	exports: [],
	declarations: [%sComponent],
	providers: [hookNavigator(%sNavigationFactory)]
})
export class %sModule{}
		`, utils.AngularString(compName), compName, utils.AngularString(compName), compName, compName, utils.AngularString(compName), utils.AngularString(compName), utils.AngularString(compName), utils.AngularString(compName)))

	}

	//Create service
	if service_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.service.ts", compName, compName)); err != nil {
		return err.Error()
	} else {
		defer service_ts_file.Close()

		service_ts_file.WriteString(fmt.Sprintf(`
import {Injectable, resolveForwardRef} from '@angular/core'
import {Subject} from 'rxjs'
		
@Injectable()
export class %sService {
	constructor() {}
}
		`, utils.AngularString(compName)))
	}

	return "Component successfully created!"
}

func NewWidget() string {
	var widgetName string
	fmt.Print(("Enter your widget name: "))
	if _, err := fmt.Scanln(&widgetName); err != nil {
		return err.Error()
	}

	os.Mkdir("src", 0755)

	//Create component folder
	if err := os.Mkdir("src/"+widgetName, 0755); err != nil {
		return err.Error()
	}

	//Create template html
	if template_html_file, err := os.Create(fmt.Sprintf("src/%s/%s.component.html", widgetName, widgetName)); err != nil {
		return err.Error()
	} else {
		defer template_html_file.Close()

		template_html_file.WriteString(fmt.Sprintf(`
<div class="p-1-16 p-r-16">
	<h1>{{'%s'}}</h1>
</div>
	`, utils.AngularString(widgetName)))
	}

	//Create style css
	if _, err := os.Create(fmt.Sprintf("src/%s/%s.component.css", widgetName, widgetName)); err != nil {
		return err.Error()
	}

	//Create component ts
	if comp_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.component.ts", widgetName, widgetName)); err != nil {
		return err.Error()
	} else {
		defer comp_ts_file.Close()

		comp_ts_file.WriteString(fmt.Sprintf(`
import {Component,Input,OnInit} from '@angular/core'

@Component({
	selector: '%s',
	templateUrl: '%s.component.html',
	styleUrls: ['%s.component.css']
})
export class %sWidgetComponent implements OnInit {
	@Input() config: {device: {id: string, name: string}}

	constructor() {}

	ngOnInit() {}
}
		`, widgetName, widgetName, widgetName, utils.AngularString(widgetName)))
	}

	//Create model ts
	if _, err := os.Create(fmt.Sprintf("src/%s/%s.model.ts", widgetName, widgetName)); err != nil {
		return err.Error()
	}

	//Create module ts
	if module_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.module.ts", widgetName, widgetName)); err != nil {
		return err.Error()
	} else {
		defer module_ts_file.Close()

		module_ts_file.WriteString(fmt.Sprintf(`
import {NgModule} from '@angular/core'
import {CoreModule, hookComponent} from '@c8y/ngx-components'
import {ContextWidgetConfig} from '@c8y/ngx-components/context-dashboard'

import {%sWidgetComponent} from './%s.component'

@NgModule({
	imports: [CoreModule],
	exports: [],
	declarations: [%sWidgetComponent],
	providers: [hookComponent({
		id: '%s.widget',
		label: '%s',
		description: '%s',
		component: %sWidgetComponent,

		data: {
			settings: {
				noNewWidgets: false,
				ng1: {
					options: {
						noDeviceTarget: false,
						groupSelectable: false
					}
				}
			}
		} as ContextWidgetConfig
	})]
})
export class %sWidgetModule{}
		`, utils.AngularString(widgetName), widgetName, utils.AngularString(widgetName), widgetName, utils.AngularString(widgetName), utils.AngularString(widgetName), utils.AngularString(widgetName), utils.AngularString(widgetName)))
	}

	//Create service ts
	if service_ts_file, err := os.Create(fmt.Sprintf("src/%s/%s.service.ts", widgetName, widgetName)); err != nil {
		return err.Error()
	} else {
		defer service_ts_file.Close()

		service_ts_file.WriteString(fmt.Sprintf(`
import {Injectable, resolveForwardRef} from '@angular/core'
import {Subject} from 'rxjs'
				
@Injectable()
export class %sWidgetService {
	constructor() {}
}
		`, utils.AngularString(widgetName)))
	}

	return "Widget successfully created!"
}
