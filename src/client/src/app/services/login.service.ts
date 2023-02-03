import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from 'src/environments/environment';
import { HashService } from 'src/app/services/hash.service';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class LoginService {

  constructor(private http:HttpClient,private router: Router) { }

  public async sendLoginRequest(username: string, password: string){
      var hashString = await HashService.getHash(password);
      var res = await firstValueFrom(this.http.post(environment.API_URL + "/logIn", {username: username, password: hashString}, {withCredentials: true}))
      if(res){
        this.router.navigate(['/'])
      }
  }
}
