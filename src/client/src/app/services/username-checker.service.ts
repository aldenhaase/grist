import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class UsernameCheckerService {
  username = "";
  constructor(private http:HttpClient) { }

  public checkUsername(username: string):Observable<boolean>{
    this.username = username;
    return this.checkIfValid() || this.checkIfTaken();
  }

  private checkIfValid(){
    if(this.username.length < 1){
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
