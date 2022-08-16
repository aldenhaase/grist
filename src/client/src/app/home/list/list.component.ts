import { Component, Inject, OnDestroy, OnInit, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable, timer, Subscription, Subject } from 'rxjs';
import { switchMap, tap, delay, share, retry, takeUntil } from 'rxjs/operators';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';


import { Router } from '@angular/router';

export interface DialogData {
  listName: string
}

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit, OnDestroy {


  private changes: Observable<boolean>;

  private stopPolling = new Subject<void>();


  public currentList:string = ""

  public isPolling:boolean = false
  
  public hideAddItem:boolean = false;

  public hideAddList:boolean = false;
  
  public AddorSub:string = "add"

  public AddorSubList:string ="add"

  public NewListName:string=""
  
  public listToDelete:string=""

  public activeTab:number = 0

  public userListArray:Array<string> = []

  public syncUserListArray:Array<string> = []

  public LocalList:User_List = {
    Last_Modified: "",
    Items: []
  };
  public SyncList:User_List={
    Last_Modified: "",
    Items: []
  }
  public NewListItem:New_Item = {
    Item: "",
    List_Name: ""
  }

  private platformId:string



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
        this.SyncList.Last_Modified = data?.Last_Modified || ""
    },
    complete:() =>{
      this.switchLists()
      this.clearField();
      if (!this.isPolling){
        this.isPolling=true
        this.startPolling()
      }
    },
      error: (error: any) => console.log(error)
    }

    private syncObserver = {
      next: (data: User_List) => {
        this.SyncList.Items = data?.Items || [];
        this.SyncList.Last_Modified = data?.Last_Modified || ""
    },
    complete:() =>{
      this.mergeLists()
    },
      error: (error: any) => console.log(error)
    }


    private setObserver = {
      next: (data: User_List) =>{
        this.SyncList.Items = data?.Items || [];
        this.SyncList.Last_Modified = data?.Last_Modified || ""
      },
      complete: () => {
        this.mergeLists();
        this.clearField();

      },
      error: (error: any) => console.log(error),
    }

    private addListObserver = {
      complete: () => {
        this.addRequest()
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
      }
    }


    private addEnumerateObserver = {
      next: (data: string[]) =>{
        this.userListArray = data || [];
            }, 
      error: (error: any) => console.log(error),
      complete: () =>{
        this.getList()
      }
    }
    private deleteListObserver = {
      error: (error: any) => console.log(error),
    complete: () => {
        this.listToDelete = ""
      }
    }

    private changesObserver = {
      next:(data: boolean) => {
        if (data){
          this.syncList()
        }
      },
      error: (error: any) => console.log(error),
    }

    private listArrayChangeObserver = {
      next:(data: boolean) => {
        if (data){
          this.enumMerge()
        }
      },
      error: (error: any) => console.log(error),
    }

    private enumMergeObserver ={
      next: (data: string[]) =>{
        this.syncUserListArray = data || [];
            }, 
      error: (error: any) => console.log(error),
      complete: () =>{
        this.mergeListNameArray()
      }
    }


    private deleteObserver = {
      next: (data: User_List) =>{
        this.SyncList.Items = data?.Items || [];
        this.SyncList.Last_Modified = data?.Last_Modified || ""
      },
      complete: () => {
        this.mergeLists();
      },
      error: (error: any) => console.log(error),
    }
  constructor(public dialog: MatDialog, private http: HttpClient, private router: Router, @Inject(PLATFORM_ID) platformId:string) {
    this.platformId = platformId
    this.changes = timer(1, 3000).pipe(
      switchMap(() => http.post<boolean>(environment.API_URL + '/checkForUpdates',  {List_Name: this.currentList, Last_Modified: this.LocalList.Last_Modified, List_Array: this.userListArray},{withCredentials: true})),
      retry(10),
      share(),
      takeUntil(this.stopPolling)
   );

   }

   openDialog(): void {
    const dialogRef = this.dialog.open(AddListDialog, {
      width: '250px',
      data:{listName: ""}
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result){
      this.NewListName = result
      this.onSubmitList()
      }
    });
  }

  openDeleteDialog(): void {
    const dialogRef = this.dialog.open(DeleteListDialog, {
      width: '250px',
      data:{listName: this.currentList}
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result){
        this.onDeleteList()
      }
    });
  }

  ngOnInit(): void {
    const getUserListReq = this.http.get<User_List>(environment.API_URL + '/checkAuth', {withCredentials: true});
    getUserListReq.subscribe(this.initObserver);
  }

  public startPolling(){
    if(isPlatformBrowser(this.platformId)){
      this.changes.subscribe(this.changesObserver)
    }
  }

  public getList(){
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/getUserList', JSON.stringify(this.currentList), {withCredentials: true});
    getUserListReq.subscribe(this.getObserver);
  }

  public enumMerge(){
    const getUserListReq = this.http.get<string[]>(environment.API_URL + '/enumerateLists', {withCredentials: true});
    getUserListReq.subscribe(this.enumMergeObserver);
  }
  public syncList() {
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/getUserList', JSON.stringify(this.currentList), {withCredentials: true});
    getUserListReq.subscribe(this.syncObserver);
  }

  public syncListArray(){
   const sync = this.http.post<boolean>(environment.API_URL + '/checkListArray',  JSON.stringify(this.userListArray),{withCredentials: true})
   sync.subscribe(this.listArrayChangeObserver)
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

  public addRequest(){
    const getUserListReq = this.http.get<string[]>(environment.API_URL + '/enumerateLists', {withCredentials: true});
    getUserListReq.subscribe(this.addEnumerateObserver);
  }

  public addNewList(){
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/createUserList', JSON.stringify(this.NewListName), {withCredentials: true});
    getUserListReq.subscribe(this.addListObserver);
  }

public deleteList(){
  const getUserListReq = this.http.post<User_List>(environment.API_URL + '/deleteUserList', JSON.stringify(this.listToDelete), {withCredentials: true});
  getUserListReq.subscribe(this.deleteListObserver);
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
    })
    var index = this.LocalList.Items.length -1
    while(index != -1){
      if(!this.SyncList.Items.includes(this.LocalList.Items[index])){
        this.LocalList.Items.splice(index,1)
   }
   index -= 1
    }
    this.LocalList.Last_Modified = this.SyncList.Last_Modified
  }

  public mergeListNameArray(){
    if(!this.syncUserListArray.includes(this.currentList)){
      alert("This List Has Been Deleted")
      this.swapList(this.syncUserListArray[0])
    }

    this.syncUserListArray.forEach(item => {
      if (!this.userListArray.includes(item)){
        this.userListArray.push(item)
      }
    })

    var index = this.userListArray.length -1
    while(index != -1){
      if(!this.syncUserListArray.includes(this.userListArray[index])){
        this.userListArray.splice(index,1)
   }
   index -= 1
    }
  }

  public setCurrentList(listName:string){
    this.currentList = listName
  }

  public swapList(listName:string){
    this.setCurrentList(listName)
    this.getList()
  }

  public switchLists(){
      this.LocalList.Items = []
      this.SyncList.Items.forEach((item, index) => {
        this.LocalList.Items.push(item)
      })
      this.LocalList.Last_Modified = this.SyncList.Last_Modified
  }

  public clearField(){
    this.NewListItem.Item = ""
  }


  public onSubmit(){
    if(this.LocalList.Items.includes(this.NewListItem.Item)){
      return
    }
    this.LocalList.Items.push(this.NewListItem.Item)
    this.addListItem()
  }


  public onSubmitList(){
    this.currentList = this.NewListName
    this.userListArray.push(this.NewListName)
    this.activeTab =  this.userListArray.indexOf(this.NewListName)
    this.addNewList()
    this.hideAddList = false
    this.AddorSubList = "add"
  }


    public onDelete(){
      let copy = this.LocalList.Items.slice()
      this.SelectedListItems.Items.forEach(item =>{
          copy.splice(copy.indexOf(item),1)
      })
      this.LocalList.Items=copy
      this.deleteListItem()
  }

  public onDeleteList(){
    var index:number
    if (this.userListArray.length === 2){
      index = 0
    }else{
      index = (this.userListArray.indexOf(this.currentList)) % (this.userListArray.length - 1)
    }
    this.listToDelete = this.currentList
    this.userListArray.splice(this.userListArray.indexOf(this.listToDelete), 1)
    this.currentList = this.userListArray[index]
    this.activeTab = this.userListArray.indexOf(this.currentList)
      this.deleteList();
      this.getList();
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

  ngOnDestroy() {
    this.stopPolling.next();
 }

}

interface User_List {
  Last_Modified: string
  Items: Array<string>
}


interface New_Item{
  Item: string
  List_Name: string
}


interface Delete_Item{
  Items: string[]
  List_Name: string
}


@Component({
  selector: 'add-list-dialog',
  templateUrl: 'add-user-list-dialog.html',
})
export class AddListDialog {
  constructor(
    public dialogRef: MatDialogRef<AddListDialog>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}

@Component({
  selector: 'delete-list-dialog',
  templateUrl: 'delete-list-dialog.html',
})
export class DeleteListDialog {
  constructor(
    public dialogRef: MatDialogRef<DeleteListDialog>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}