import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable, firstValueFrom } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class UsernameCheckerService {
  username = "";
  constructor(private http:HttpClient) { }

  async checkUsername(username: string){
    this.username = username;
    var res = await firstValueFrom(this.checkIfValid() || this.checkIfTaken());
    return res
  }

  private checkIfValid(){
    if(this.username.length == 0){
      return new Observable<boolean>(observer => {
        observer.next( true )
        observer.complete()
     });
    }
    return null;
  }

  private checkIfTaken(){
    return this.http.post<boolean>(environment.API_URL + "/checkUsername", {username: this.username})
  }
}
