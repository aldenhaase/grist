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
  userValid: boolean = true;
  userReason: string = '';
  passValid: boolean = true;
  passReason: string = '';
  constructor(private http: HttpClient) { }

  public onSubmit() {
    if(this.password != "hello"){
      this.passValid = false;
    }
    if(this.username != "hello"){
      this.userValid = false;
    }
    console.log(this.password);
    console.log(this.username);
  }

  public checkUsername(event: string){
    this.sendAPI();
  }
  public checkPassword(event: string){
    if(event != "hello"){
      this.passValid = false;
    }else{
      this.passValid = true;
    }
  }
  private sendAPI(){
    this.http.post<usernameResponse>(environment.API_URL + "/checkUsername", {username: this.username}).subscribe(data => {
      if(!data.Valid){
        this.userValid = false;
      }
  })
  }


}

interface usernameResponse{
    Valid: boolean;
    reason: string;
}
