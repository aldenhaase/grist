import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Router } from '@angular/router';
@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {
  public currentList:string = ""
  
  public hideAddItem:boolean = false;

  public hideAddList:boolean = false;
  
  public AddorSub:string = "add"

  public AddorSubList:string ="add"

  public NewListName:string=""

  public userListArray:Array<string> = []

  public LocalList:User_List = {
    Items: []
  };
  public SyncList:User_List={
    Items: []
  }
  public NewListItem:New_Item = {
    Item: "",
    List_Name: ""
  }



  public SelectedListItems:Delete_Item = {
    Items: [],
    List_Name: ""
  }

    private initObserver = {
    complete:() =>{
      this.enumerateListsRequest();
    },
      error: (error: any) => this.router.navigate(['/login'])
    }


    private getObserver = {
      next: (data: User_List) => {
        this.SyncList.Items = data?.Items || [];
    },
    complete:() =>{
      this.switchLists()
      this.clearField();

    },
      error: (error: any) => console.log(error)
    }
    private setObserver = {
      next: (data: User_List) =>{
        this.SyncList.Items = data?.Items || [];
      },
      complete: () => {
        this.mergeLists();
        this.clearField();

      },
      error: (error: any) => console.log(error),
    }

    private addListObserver = {
      complete: () => {
        this.enumerateListsRequest()

      },
      error: (error: any) => console.log(error),
    }

    private enumerateObserver = {
      next: (data: string[]) =>{
        this.userListArray = data || [];
            }, 
      error: (error: any) => console.log(error),
      complete: () =>{
        this.currentList = this.userListArray[0]
        this.getList();//do something with undefined here
        this.currentList = this.userListArray[0]
      }
    }


    private deleteObserver = {
      next: (data: User_List) =>{
        this.SyncList.Items = data?.Items || [];
      },
      complete: () => {
        this.mergeLists();
      },
      error: (error: any) => console.log(error),
    }
  constructor(private http: HttpClient, private router: Router) { }

  ngOnInit(): void {
    const getUserListReq = this.http.get<User_List>(environment.API_URL + '/checkAuth', {withCredentials: true});
    getUserListReq.subscribe(this.initObserver);
  }

  public getList(){
    
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/getUserList', JSON.stringify(this.currentList), {withCredentials: true});
    getUserListReq.subscribe(this.getObserver);
  }

  public addListItem(){
    this.NewListItem.List_Name = this.currentList
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/setUserList', this.NewListItem, {withCredentials: true});
    getUserListReq.subscribe(this.setObserver);
  }



  public enumerateListsRequest(){
    const getUserListReq = this.http.get<string[]>(environment.API_URL + '/enumerateLists', {withCredentials: true});
    getUserListReq.subscribe(this.enumerateObserver);
  }

  public addNewList(){
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/createUserList', JSON.stringify(this.NewListName), {withCredentials: true});
    getUserListReq.subscribe(this.addListObserver);
  }

  public deleteListItem(){
    this.SelectedListItems.List_Name = this.currentList
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/deleteListItem', this.SelectedListItems, {withCredentials: true});
    getUserListReq.subscribe(this.deleteObserver);
  }

  public mergeLists(){
    this.SyncList.Items.forEach(item => {
      if (!this.LocalList.Items.includes(item)){
        this.LocalList.Items.push(item)
      }
    });
    this.LocalList.Items.forEach((item,index) => {
      if(!this.SyncList.Items.includes(item)){
            this.LocalList.Items.splice(index,1)
      }
    });
  }

  public setCurrentList(listName:string){
    this.currentList = listName
  }

  public swapList(listName:string){
    this.setCurrentList(listName)
    this.getList()
  }

  public switchLists(){
    if (this.SyncList.Items.length == 0){
      this.LocalList.Items = []
    }
      this.SyncList.Items.forEach((item, index) => {
        this.LocalList.Items[index] = item
      })
  }

  public clearField(){
    this.NewListItem.Item = ""
  }


  public onSubmit(){
    this.LocalList.Items.push(this.NewListItem.Item)
    this.addListItem()
  }


  public onSubmitList(){
    this.addNewList()
  }


    public onDelete(){
      let copy = this.LocalList.Items.slice()
      this.SelectedListItems.Items.forEach(item =>{
          copy.splice(copy.indexOf(item),1)
      })
      this.LocalList.Items=copy
      this.deleteListItem()
  }
  public toggleHideAddItem(){
    this.hideAddItem = !this.hideAddItem;
    if (this.hideAddItem){
      this.AddorSub ="minimize"
    }else{
      this.AddorSub = "add"
    }
  }

  public toggleHideAddList(){
    this.hideAddList= !this.hideAddList;
    if (this.hideAddList){
      this.AddorSubList ="minimize"
    }else{
      this.AddorSubList = "add"
    }
  }

}

interface User_List {
  Items: Array<string>;
}


interface New_Item{
  Item: string
  List_Name: string
}


interface Delete_Item{
  Items: string[]
  List_Name: string
}