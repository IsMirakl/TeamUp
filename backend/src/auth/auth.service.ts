import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';
import { PrismaService } from 'src/prisma/prisma.service';
import { LoginDto } from './dto/login.dto';
import { RegisterDto } from './dto/register.dto';

@Injectable()
export class AuthService {
  constructor(
    private prisma: PrismaService,
    private jwtService: JwtService,
  ) {}

  async register(dto: RegisterDto) {
    const hashedPassword = await bcrypt.hash(dto.password, 12);
    
    const user = await this.prisma.user.create({
      data: {
        email: dto.email,
        name: dto.name,
        emailVerifed: false,
        role: 'USER',
        subscriptionPlan: 'FREE',
      }
    });
    
    await this.prisma.account.create({
      data: {
        userId: user.id,
        type: 'credentials',
        Provider: 'credentials',
        passwordHash: hashedPassword,
      }
    });

    return {
      access_token: this.jwtService.sign({
        sub: user.id,
        email: user.email
      }),
      user: { id: user.id, email: user.email, name: user.name }
    };
  }

  async login(dto: LoginDto) {
    const account = await this.prisma.account.findFirst({
      where: {
        type: 'credentials',
        Provider: 'credentials',
      },
      include: { user: true }
    });

    if (!account || !await bcrypt.compare(dto.password, account.passwordHash)) {
      throw new UnauthorizedException('Неверный email/пароль');
    }

    return {
      access_token: this.jwtService.sign({
        sub: account.user.id,
        email: account.user.email
      }),
      user: { 
        id: account.user.id, 
        email: account.user.email, 
        name: account.user.name 
      }
    };
  }
}
