<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\View\View;
use \Illuminate\Http\JsonResponse;

class UserController extends Controller
{
    /**
     * Show the profile for a given user.
     */
    public function show(string $id): JsonResponse
    {
        return response()->json([
            'user' => User::findOrFail($id)
        ]);
    }
}
