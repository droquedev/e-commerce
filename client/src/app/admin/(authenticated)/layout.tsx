import { ChildrenProps } from '@/types'
import React from 'react'

export default function layout({children}: Readonly<ChildrenProps>) {
  return (
    <div>
      layout
      {children}
    </div>
  )
}
