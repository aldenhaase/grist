import { TestBed } from '@angular/core/testing';

import { ListGrabberService } from './list-grabber.service';

describe('ListGrabberService', () => {
  let service: ListGrabberService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ListGrabberService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
