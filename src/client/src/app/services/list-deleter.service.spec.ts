import { TestBed } from '@angular/core/testing';

import { ListDeleterService } from './list-deleter.service';

describe('ListDeleterService', () => {
  let service: ListDeleterService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ListDeleterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
