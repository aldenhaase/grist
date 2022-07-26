import { Component,} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
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
  constructor(private http: HttpClient) { }

  public async onSubmit() {
    var hash = await this.digest(this.username)
    var hashString = this.digestToString(hash)
    this.http.post<registerUserResponse>(environment.API_URL + "/registerNewUser", {username: this.username, password: hashString}).subscribe(data => {
     if(data.status) {
      console.log("error registering new user")
     }
    })
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
               this.http.post<usernameResponse>(environment.API_URL + "/checkUsername", {username: this.username}).subscribe(data => {
        this.userValid = !data.exists;
        this.userReason = data.reason;
  })
  }
  public checkPassword(){
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
