import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class PasswordCheckerService {
  password = "";
  constructor(private http:HttpClient) { }

  public checkPassword(password: string):Observable<boolean>{
    this.password = password;
    return this.checkIfValid();
  }

  private checkIfValid(){
    if(this.password.length < 3){
      return new Observable<boolean>(observer => {
        observer.next( false )
        observer.complete()
     });
    }
    return new Observable<boolean>(observer => {
      observer.next( true )
      observer.complete()
   });
  }

}
