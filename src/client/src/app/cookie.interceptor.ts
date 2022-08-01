import { ContentChild, Inject, Injectable, Optional } from '@angular/core';
import { Request } from 'express';
import { REQUEST } from '@nguniversal/express-engine/tokens'
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
  HttpHeaders
} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class CookieInterceptor implements HttpInterceptor {

  constructor(
    @Optional() @Inject(REQUEST) protected serverReq: Request
  ) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    let cookie = this.serverReq.headers.cookie || ""
    let headers = new HttpHeaders()
    .set('content-type', 'application/json')
    .set('Cookie', cookie);
    const authReq = request.clone({headers: headers});
    return next.handle(authReq);
  }
}
