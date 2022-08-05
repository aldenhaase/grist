import { Component,} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Router } from '@angular/router';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent{
  password: string = '';
  username: string = '';
  userValid: boolean = false;
  userReason: string = '';
  passValid: boolean = false;
  passReason: string = '';
  constructor(private http: HttpClient, private router: Router) { }
    private observer = {
      error: (error: any) => console.log("failed register new user"),
      complete: () => this.router.navigate(['/login']),
    }

    private getCookieObserver = {
      error: (error: any) => console.log("failed to get registration cookie"),
      complete: async () => {
        var hash = await this.digest(this.password)
        var hashString = this.digestToString(hash)
        var registerReq = this.http.post<registerUserResponse>(environment.API_URL + "/registerNewUser", {username: this.username, password: hashString}, {withCredentials: true});
        registerReq.subscribe(this.observer)
      }
    }

  public async onSubmit() {
    var getRegCookies = this.http.get<registerUserResponse>(environment.API_URL + "/getRegistrationCookies", {withCredentials: true});
    getRegCookies.subscribe(this.getCookieObserver)
  }
  private digestToString(buffer: ArrayBuffer){
    const byteArray = new Uint8Array(buffer);

    const hexCodes = [...byteArray].map(value => {
      const hexCode = value.toString(16);
      const paddedHexCode = hexCode.padStart(2, '0');
      return paddedHexCode;
    });
  
    return hexCodes.join('');
  }

  private async digest(text: string){
    var encoder = new TextEncoder();
    var encoded = encoder.encode(this.username.normalize())
    var hash = await crypto.subtle.digest("SHA-256", encoded);
    return hash
  }
  public checkUsername(){      
    if(this.username.length < 1){
      this.userValid = false;
      this.userReason = "username must exist"
      return
    }
        this.http.post<usernameResponse>(environment.API_URL + "/checkUsername", {username: this.username}).subscribe(data => {
        this.userValid = !data.exists;
        this.userReason = data.reason;
  })
  }
  public checkPassword(){
    if(this.password.length < 1){
      this.passValid = false;
      this.passReason = "password must be longer than 1"
      return
    }
   this.passValid = true; 
  }


}
interface registerUserResponse{
  status: number;
  error:  string;
}
interface usernameResponse{
    exists: boolean;
    reason: string;
}
