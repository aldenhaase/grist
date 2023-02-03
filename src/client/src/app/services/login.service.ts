import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from 'src/environments/environment';
import { HashService } from 'src/app/services/hash.service';

@Injectable({
  providedIn: 'root'
})

export class LoginService {

  constructor(private http:HttpClient,private router: Router) { }

  private observer = {
    error: (error: any) => console.log(error),
    complete: () => this.router.navigate(['/'])
  }

  public async sendLoginRequest(username: string, password: string){
      var hashString = await HashService.getHash(password);
      var loginReq = this.http.post(environment.API_URL + "/logIn", {username: username, password: hashString}, {withCredentials: true});
      loginReq.subscribe(this.observer);
  }
}
