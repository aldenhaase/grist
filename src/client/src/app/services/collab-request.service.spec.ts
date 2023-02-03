import { TestBed } from '@angular/core/testing';

import { CollabRequestService } from './collab-request.service';

describe('CollabRequestService', () => {
  let service: CollabRequestService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CollabRequestService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
