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
    private observer = {
      next: (data: userList) => this.list = data,
      error: (error: any) => this.router.navigate(['/login'])
    }

  constructor(private http: HttpClient, private router: Router) { }
  list = {} as userList;
  ngOnInit(): void {
    const getUserListReq = this.http.get<userList>(environment.API_URL + '/getUserList', {withCredentials: true});
    getUserListReq.subscribe(this.observer);
  }


}

interface userList {
  listName: string;
  listItems: Array<string>;
}
