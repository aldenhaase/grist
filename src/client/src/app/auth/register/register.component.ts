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

  public onSubmit() {
    this.http.post<usernameResponse>(environment.API_URL + "/setUsername", {username: this.username}).subscribe(data => {
      this.userValid = data.exists;
      this.userReason = data.reason;
})
  }

  public checkUsername(){
        this.http.post<usernameResponse>(environment.API_URL + "/checkUsername", {username: this.username}).subscribe(data => {
        this.userValid = !data.exists;
        this.userReason = data.reason;
  })
  }
  public checkPassword(){
    this.passValid = true
  }


}

interface usernameResponse{
    exists: boolean;
    reason: string;
}
