import { TestBed } from '@angular/core/testing';

import { UsernameCheckerService } from './username-checker.service';

describe('UsernameCheckerService', () => {
  let service: UsernameCheckerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UsernameCheckerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
