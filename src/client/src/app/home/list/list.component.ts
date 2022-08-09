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
        this.SyncedList.Title = data.Title || "";
        this.SyncedList.Items = data.Items || [];
    },
      error: (error: any) => this.router.navigate(['/login'])
    }

    private getObserver ={
      next: (data: User_List) => {
        this.SyncedList.Title = data.Title || "";
        this.SyncedList.Items = data.Items || [];
      },
      error: (error: any) => console.log(error),
      complete: () => this.clearField()
    }

    private setObserver = {
      complete: () => this.updateList(),
      error: (error: any) => console.log(error),
    }

  constructor(private http: HttpClient, private router: Router) { }
  public LocalList:User_List = {
    Title: "",
    Items: []
  };
  public SyncedList:User_List={
    Title: "",
    Items: []
  }
  newListItem = ""
  ngOnInit(): void {
    const getUserListReq = this.http.get<User_List>(environment.API_URL + '/getUserList', {withCredentials: true});
    getUserListReq.subscribe(this.initObserver);
  }

  public addListItem(){
    const getUserListReq = this.http.post<User_List>(environment.API_URL + '/setUserList', this.LocalList, {withCredentials: true});
    getUserListReq.subscribe(this.setObserver);
  }

  public clearField(){
    this.newListItem = ""
  }

  public updateList(){
    const getUserListReq = this.http.get<User_List>(environment.API_URL + '/getUserList', {withCredentials: true});
    getUserListReq.subscribe(this.getObserver);
  }

  public onSubmit(){
    this.LocalList.Items.push(this.newListItem)
    this.addListItem()
  }
}

interface User_List {
  Title: string;
  Items: Array<string>;
}

