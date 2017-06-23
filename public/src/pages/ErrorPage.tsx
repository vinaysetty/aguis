import * as React from "react";
import { ViewPage } from "./ViewPage";

class ErrorPage extends ViewPage {
    private pages: { [key: string]: JSX.Element } = {};
    constructor() {
        super();
        this.defaultPage = "404";
        this.pages["404"] = <div><h1>404 Page not found</h1><p>The page you where looking for does not exist</p></div>;
    }

    public pageNavigation(page: string): void {
        "not implemented";
    }

    public renderContent(page: string): JSX.Element {
        if (page.length === 0) {
            page = this.defaultPage;
        }
        return <div><h1>404 Page not found</h1><p>The page you where looking for does not exist</p></div>;
    }
}

export { ErrorPage };
