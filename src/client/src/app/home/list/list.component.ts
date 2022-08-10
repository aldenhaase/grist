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
    private initObserver = {
      next: (data: User_List) => {
        this.LocalList.Title = data.Title || "";
        this.LocalList.Items = data.Items || [];
    },
      error: (error: any) => this.router.navigate(['/login'])
    }

    private setObserver = {
      next: (data: User_List) =>{
        this.SyncList.Title = data.Title || "";
        this.SyncList.Items = data.Items || [];
      },
      complete: () => this.mergeLists(),
      error: (error: any) => console.log(error),
    }

  constructor(private http: HttpClient, private router: Router) { }
  public LocalList:User_List = {
    Title: "",
    Items: []
  };
  public SyncList:User_List={
    Title: "",
    Items: []
  }
  public NewListItem:New_Item = {
    Item: ""
  }
  ngOnInit(): void {
    const getUserListReq = this.http.get<User_List>(environment.API_URL + '/getUserList', {withCredentials: true});
    getUserListReq.subscribe(this.initObserver);
  }

  public addListItem(){
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/setUserList', this.NewListItem, {withCredentials: true});
    getUserListReq.subscribe(this.setObserver);
  }

  public mergeLists(){
    this.SyncList.Items.forEach(item => {
      if (!this.LocalList.Items.includes(item)){
        this.LocalList.Items.push(item)
      }
    });
    this.clearField()
  }

  public clearField(){
    this.NewListItem.Item = ""
  }


  public onSubmit(){
    this.LocalList.Items.push(this.NewListItem.Item)
    this.addListItem()
  }
}

interface User_List {
  Title: string;
  Items: Array<string>;
}

interface New_Item{
  Item: string
}