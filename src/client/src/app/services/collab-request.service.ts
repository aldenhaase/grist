import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { collection } from '../types/listTypes';
import { collaboratorQuery } from '../types/collaboratorTypes';
import { Observable, firstValueFrom} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CollabRequestService {

  constructor(private http:HttpClient) { }
  async addCollaborator(listName: string, shareWith: string){
    var collaboratorQuery = {listName: listName, shareWith: shareWith}
    await firstValueFrom(this.http.post<collaboratorQuery>(environment.API_URL + "/collaborator", JSON.stringify(collaboratorQuery), {withCredentials: true}))
  }
}
