import * as React from "react";
import { ILink, NavigationManager } from "../managers";
import { View, ViewPage } from "./ViewPage";
import { HelpView } from "./views/HelpView";

import { INavInfo, NavigationHelper } from "../NavigationHelper";

export class HelpPage extends ViewPage {

    private navMan: NavigationManager;
    private pages: { [name: string]: JSX.Element } = {};

    constructor(navMan: NavigationManager) {
        super();
        this.template = "frontpage";
        this.navMan = navMan;
        this.navHelper.defaultPage = "help";
        this.navHelper.registerFunction("help", this.help);
    }

    public async help(info: INavInfo<any>): View {
        return <HelpView></HelpView>;
    }
}