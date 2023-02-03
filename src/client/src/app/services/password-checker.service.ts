import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable, firstValueFrom} from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class PasswordCheckerService {
  constructor(private http:HttpClient) { }

  async checkPassword(password: string){
    return await firstValueFrom(this.checkIfValid(password))
  }

  private checkIfValid(password: string){
    if(password.length < 3){
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
