/*
 * Copyright (c) 2021, Fraunhofer AISEC. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *                    $$$$$$\  $$$$$$$\   $$$$$$\
 *                   $$  __$$\ $$  __$$\ $$  __$$\
 *                   $$ /  \__|$$ |  $$ |$$ /  \__|
 *                   $$ |      $$$$$$$  |$$ |$$$$\
 *                   $$ |      $$  ____/ $$ |\_$$ |
 *                   $$ |  $$\ $$ |      $$ |  $$ |
 *                   \$$$$$   |$$ |      \$$$$$   |
 *                    \______/ \__|       \______/
 *
 */
package cpg

import (
	"log"

	"tekao.net/jnigi"
)

func NewString(s string) *jnigi.ObjectRef {
	o, err := env.NewObject("java/lang/String", []byte(s))
	if err != nil {
		log.Fatal(err)

	}

	return o
}

func NewCharSequence(s string) *jnigi.ObjectRef {
	o, err := env.NewObject("java/lang/String", []byte(s))
	if err != nil {
		log.Fatal(err)

	}

	return o.Cast("java/lang/CharSequence")
}

func NewBoolean(b bool) *jnigi.ObjectRef {
	// TODO: Use Boolean.valueOf
	o, err := env.NewObject("java/lang/Boolean", b)
	if err != nil {
		log.Fatal(err)
	}

	return o
}

func NewInteger(i int) *jnigi.ObjectRef {
	// TODO: Use Integer.valueOf
	o, err := env.NewObject("java/lang/Integer", i)
	if err != nil {
		log.Fatal(err)
	}

	return o
}

func NewDouble(d float64) *jnigi.ObjectRef {
	// TODO: Use Integer.valueOf
	o, err := env.NewObject("java/lang/Double", d)
	if err != nil {
		log.Fatal(err)
	}

	return o
}
