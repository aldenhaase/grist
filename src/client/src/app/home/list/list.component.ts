import { Component, OnInit, AfterViewInit, ViewChild, Inject} from '@angular/core';
import { Observable, firstValueFrom, interval } from 'rxjs';
import { collection, item, list} from 'src/app/types/listTypes';
import { ListGrabberService } from 'src/app/services/list-grabber.service';
import { ListSetterService } from 'src/app/services/list-setter.service';
import { MergeService } from 'src/app/services/merge.service';
import { collaboratorQuery } from 'src/app/types/collaboratorTypes';
import { CollabRequestService } from 'src/app/services/collab-request.service';
import { ListDeleterService } from 'src/app/services/list-deleter.service';
import { ItemDeleterService } from 'src/app/services/item-deleter.service';
import {MatPaginator} from '@angular/material/paginator';
import {MatTableDataSource} from '@angular/material/table';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { ListDialogData, CollaboratorDialogData } from 'src/app/types/dialogTypes';
import { UsernameCheckerService } from 'src/app/services/username-checker.service';




@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit{
  constructor(private listGrabberService: ListGrabberService,
              private listSetterService: ListSetterService,
              private mergeService: MergeService,
              private collabRequestService: CollabRequestService,
              private listDeleterService: ListDeleterService,
              private itemDeleterService: ItemDeleterService,
              public dialog: MatDialog,
              ){}
  localCollection:collection = {lists: []}
  delCollection:collection = {lists: []}
  remoteCollection$ = new Observable<collection>;
  setCollection$ = new Observable<collection>;
  addCollab$ = new Observable<collaboratorQuery>;
  deleteList$ = new Observable<list>;
  deleteItem$ = new Observable<item>;
  activeList = 0;
  selectedItems:item[] = [];
  syncTimer: ReturnType<typeof setTimeout> = this.setSyncTimer()

  collaborator = "";
  collabList = "";
  newValue = "";

  async ngOnInit(){
    this.remoteCollection$ = this.listGrabberService.grab()
    await this.mergeService.merge(this.remoteCollection$, this.localCollection)
  }

  setSyncTimer(){
    return setInterval(() => { 
      this.mergeService.merge(this.remoteCollection$, this.localCollection); 
  }, 2000);
  }
  async addItem(){
    var newItem = {value: this.newValue, marked: false, uuid: self.crypto.randomUUID()};
    this.localCollection.lists[this.activeList].items = this.localCollection.lists[this.activeList].items?.concat(newItem) || [newItem]
    this.newValue = ""
    this.setCollection()
  }
   async addCollaberator(collaborator: string){
    clearInterval(this.syncTimer)
    await this.collabRequestService.addCollaborator(this.localCollection.lists[this.activeList].listName, collaborator)
    this.syncTimer = this.setSyncTimer()
  }

  async deleteList(){
    clearInterval(this.syncTimer)
    var listToDelete = this.localCollection.lists[this.activeList]
    this.localCollection.lists.splice(this.activeList,1)
    await this.listDeleterService.delete(listToDelete)
    this.syncTimer = this.setSyncTimer()
  }
  async deleteSelected(){
    clearInterval(this.syncTimer)
    var itemsToDelete:item[] = this.copy(this.localCollection.lists[this.activeList].items.filter(function (item) {
      return item.marked == true;
    }))
    this.localCollection.lists[this.activeList].items = this.localCollection.lists[this.activeList].items.filter(function (item) {
      return item.marked == false;
    })

    await this.itemDeleterService.delete(itemsToDelete)
    this.syncTimer = this.setSyncTimer()

  }

  async setCollection(){
    clearInterval(this.syncTimer)
   await this.listSetterService.set(this.localCollection)
   this.syncTimer = this.setSyncTimer()
  }

  addMarkAndSet(item:item){
    item.marked = !item.marked
    this.setCollection()
  }

  checkForMarkedItems(){
    var marked = false
     this.localCollection.lists.forEach(list =>{
      list.items.forEach(item =>{
        if (item.marked){
          marked = true
        }
      })
    })
    return marked
  }

  async mergeCollections(){
    clearInterval(this.syncTimer)
    await this.mergeService.merge(this.remoteCollection$, this.localCollection)
    this.syncTimer = this.setSyncTimer()

  }
  
  trackListByFn(index:any, list:list){
    return list.uuid
  }
  trackItemByFn(index:any, item:item){
    return item.uuid
  }
  
  copy(object:any){
    return JSON.parse(JSON.stringify(object))
  }
  openAddList() {
    const dialogRef = this.dialog.open(AddListDialog, {
      width: '250px',
      data:{listName: ""}
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result){
        this.localCollection.lists = this.localCollection.lists.concat({listName: result, items: [], uuid: self.crypto.randomUUID()})
        this.activeList = this.localCollection.lists.length-1
        this.setCollection()
      }
    });
  }

  openDeleteList() {
    const dialogRef = this.dialog.open(DeleteListDialog, {
      width: '250px',
      data:{listName: this.localCollection.lists[this.activeList].listName}
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result){
        this.deleteList()
      }
    });
  }

  openAddCollaborator() {
    const dialogRef = this.dialog.open(AddCollaboratorDialog, {
      width: '250px',
      data:{collaborator: "", valid: false}
    });

    dialogRef.afterClosed().subscribe(result => {
      if(result){
        this.addCollaberator(result)
      }
    });
  }
  
}


@Component({
  selector: 'add-list-dialog',
  templateUrl: 'add-user-list-dialog.html',
})
export class AddListDialog {
  constructor(
    public dialogRef: MatDialogRef<AddListDialog>,
    @Inject(MAT_DIALOG_DATA) public data: ListDialogData,
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
    @Inject(MAT_DIALOG_DATA) public data: ListDialogData,
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}

@Component({
  selector: 'add-collaborator-dialog',
  templateUrl: 'add-collaborator-dialog.html',
})
export class AddCollaboratorDialog {
  collaboratorIsValid$ = new Observable<boolean>;
  constructor(
    public dialogRef: MatDialogRef<AddCollaboratorDialog>,
    private usernameService: UsernameCheckerService, 
    @Inject(MAT_DIALOG_DATA) public data: CollaboratorDialogData,
  ) {}
  async checkUsername(username: string){
    this.data.valid = !await(this.usernameService.checkUsername(this.data.collaborator));
  }
  onNoClick(): void {
    this.dialogRef.close();
  }
}