import { ItemType as BreadItemType } from "antd/es/breadcrumb/Breadcrumb";
import { ItemType } from "antd/es/menu/hooks/useItems";
import { makeAutoObservable } from "mobx";

export default class Site {
    pageTile = 'ZUES.PUB'
    breadcrumb: Array<BreadItemType> = [{ title: 'home' }]
    menuItems: ItemType[] = []

    constructor() {
        makeAutoObservable(this);
    }

    setPageTitle(title: string) {
        this.pageTile = title
    }

    setBreadcrumb(breadcrumb: Array<BreadItemType>) {
        this.breadcrumb = breadcrumb
    }

    setMenu(items: ItemType[]) {
        this.menuItems = items
    }
}
