declare module '*.vue';
declare module "@codemirror/*"

declare type status = "success" | "error" | string
declare type Page = {
    selected: number
}
declare type TableHeader = {
    Id: string
    Name: string
    Info: string
    Maxscore: number
}
declare type TableCell = {
    Username: string
    Firstname:string
    Lastname:string
    Score:number,
    Data: {
        [key: string]: {
            id: string,
            Pass: boolean,
            Timestamp: number,
            Attempt: number,
            Score: number
        }
    }
}
declare type TableData = {
    response: {
        Structure: { [key: string]: TableHeader }
        Assigned: Array<TableCell>
    }
    status: status,
}

declare type Groupdata = {
    
    response: Array<
        {
            Assignment: {
                [key: string]: {
                    id: string,
                    sol_maxcase: boolean,
                    sol_randomcase: boolean,
                    sol_randomcasenum: number,
                    sol_randomcasenumrange: Array<number>,
                }
            }
            Groupname: string,
            Active: boolean
        }>
    status: string
}

declare type Assignment = {
    [key: string]: {
        id: string
        Pass: boolean
    }
}

declare type GroupProbs = {
    response: Array<{
            [key: string]: {
                Id: string
                Name: string,
                Info: string
            }
    }>
    status: string
}

declare type UserData = {
    response: Array<{
        Username:string
        Firstname:string
        Lastname:string
        Groups:Array<string>
    }>
    status: string
}

declare type SelfAssignment = {
    response: Array<{
        Id:string
        Name:string
        Info:string
        Link:string
        Pass:string
    }>
    status: string
}

declare type GroupAssignment = {
    response: {[key:string]: {
        Assigned:Array<{
            Id: string
            Name: string
            Info: string
            Link: string
            Testcase: Array<{name:string, stdin:boolean}>
        }>
    }}
    status: string
}