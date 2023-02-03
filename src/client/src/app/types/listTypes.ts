

export interface collection {
    lists: list[]
}

export interface list{
    listName: string
    items: item[]
    uuid: string
}

export interface item{
    value: string
    marked: boolean
    uuid: string
}
