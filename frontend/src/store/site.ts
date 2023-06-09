import { ItemType } from "antd/es/menu/hooks/useItems";
import { makeAutoObservable } from "mobx";

export default class Site {
    menuItems: ItemType[] = [{ label: '主页', key: '/' }]

    constructor() {
        makeAutoObservable(this);
    }

    setMenu(items: ItemType[]) {
        this.menuItems = items
    }
}
