import { Controller, Get, UseGuards } from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { Roles } from 'src/auth/guards/role.decorator';
import { RolesGuard } from 'src/auth/guards/roles.guard';

@Controller('profile')
export class ProfileController {
    @Get('admin')
    @UseGuards(JwtAuthGuard, RolesGuard)
    @Roles('ADMIN')
    adminonly() {
      return 'Admin access granted';
    }
}
