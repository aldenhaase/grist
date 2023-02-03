import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from 'src/environments/environment';
import { HashService } from 'src/app/services/hash.service';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {
  constructor(private http:HttpClient,private router: Router) { }

  public async sendRegisterRequest(username: string, password: string){
    var hashedPassword = await HashService.getHash(password);
    await firstValueFrom(this.http.get(environment.API_URL + "/getRegistrationCookies", {withCredentials: true}))
    var registerRes = await firstValueFrom(this.http.post(environment.API_URL + "/registerNewUser", {username: username, password: hashedPassword}, {withCredentials: true}))
    if(registerRes){
      this.router.navigate(['/login'])
    }
}
}