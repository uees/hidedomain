import { ItemType as BreadItemType } from "antd/es/breadcrumb/Breadcrumb";
import { ItemType } from "antd/es/menu/hooks/useItems";
import { makeAutoObservable } from "mobx";

export default class Site {
    breadcrumb: Array<BreadItemType> = [{ title: '主页' }]
    menuItems: ItemType[] = [{ label: '主页', key: 'home' }]

    constructor() {
        makeAutoObservable(this);
    }

    setBreadcrumb(breadcrumb: Array<BreadItemType>) {
        this.breadcrumb = breadcrumb
    }

    setMenu(items: ItemType[]) {
        this.menuItems = items
    }
}
