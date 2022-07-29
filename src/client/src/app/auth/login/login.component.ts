import { HttpClient } from '@angular/common/http';
import { Component,} from '@angular/core';
import { Router } from '@angular/router';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent{
    private observer = {
      error: (error: any) => console.log(error),
      complete: () => this.router.navigate(['/'])
    }

  password: string = '';
  username: string = '';
  loginValid: boolean = false;
  formValid: boolean = false;
  constructor(private http: HttpClient,private router: Router) { }

  public async onSubmit() {
      var hash = await this.digest(this.password)
      var hashString = this.digestToString(hash)
        var loginReq = this.http.post(environment.API_URL + "/logIn", {username: this.username, password: hashString}, {withCredentials: true});
        loginReq.subscribe(this.observer);
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

public validateForm(){
if(this.password.length < 1 || this.username.length < 1){
      this.formValid = false;
      return
    }
   this.formValid = true; 
}

}
