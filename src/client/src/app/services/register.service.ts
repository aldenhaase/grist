import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from 'src/environments/environment';
import { HashService } from 'src/app/services/hash.service';
@Injectable({
  providedIn: 'root'
})
export class RegisterService {
  username = '';
  password = '';
  constructor(private http:HttpClient,private router: Router) { }

  public async sendRegisterRequest(username: string, password: string){
    this.password = await HashService.getHash(password);
    this.username = username;
    var getRegCookies = this.http.get(environment.API_URL + "/getRegistrationCookies", {withCredentials: true});
    getRegCookies.subscribe(this.getCookieObserver)
}

private observer = {
  error: (error: any) => console.log("failed to create list"),
  complete: () => this.router.navigate(['/login']),
}
private getCookieObserver = {
  error: (error: any) => console.log("failed to get registration cookie"),
  complete: async () => {
    var registerReq = this.http.post(environment.API_URL + "/registerNewUser", {username: this.username, password: this.password}, {withCredentials: true});
    registerReq.subscribe(this.observer)
  }
}
}
